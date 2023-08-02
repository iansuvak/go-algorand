// Copyright (C) 2019-2023 Algorand, Inc.
// This file is part of go-algorand
//
// go-algorand is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// go-algorand is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with go-algorand.  If not, see <https://www.gnu.org/licenses/>.

package network

import (
	"context"
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"net/http"
	"path/filepath"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/algorand/go-algorand/config"
	"github.com/algorand/go-algorand/logging"
	"github.com/algorand/go-algorand/network/p2p"
	"github.com/algorand/go-algorand/network/p2p/peerstore"
	"github.com/algorand/go-algorand/protocol"
	"github.com/algorand/go-deadlock"
	"github.com/algorand/websocket"

	"github.com/labstack/gommon/log"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	p2ppeerstore "github.com/libp2p/go-libp2p/core/peerstore"
	"github.com/libp2p/go-libp2p/p2p/host/peerstore/pstoremem"
)

const peerStorePath = "peerstore.db"

var outgoingMessagesBufferSize = int(
	max(config.Consensus[protocol.ConsensusCurrentVersion].NumProposers,
		config.Consensus[protocol.ConsensusCurrentVersion].SoftCommitteeSize,
		config.Consensus[protocol.ConsensusCurrentVersion].CertCommitteeSize,
		config.Consensus[protocol.ConsensusCurrentVersion].NextCommitteeSize) +
		max(config.Consensus[protocol.ConsensusCurrentVersion].LateCommitteeSize,
			config.Consensus[protocol.ConsensusCurrentVersion].RedoCommitteeSize,
			config.Consensus[protocol.ConsensusCurrentVersion].DownCommitteeSize),
)

// P2PNetwork implements the GossipNode interface
type P2PNetwork struct {
	service     *p2p.Service
	log         logging.Logger
	config      config.Local
	genesisID   string
	networkID   protocol.NetworkID
	ctx         context.Context
	ctxCancel   context.CancelFunc
	peerStats   map[peer.ID]*p2pPeerStats
	peerStatsMu deadlock.Mutex

	wg sync.WaitGroup

	// which tags to use with libp2p's GossipSub, mapped to topic names
	topicTags map[protocol.Tag]string

	handlers Multiplexer

	// legacy websockets message support
	broadcaster        broadcaster
	wsReadBuffer       chan IncomingMessage
	peers              map[peer.ID]*wsPeer
	peersLock          deadlock.RWMutex
	peersChangeCounter int32
}

type p2pPeerStats struct {
	txReceived uint64
}

// NewP2PNetwork returns an instance of GossipNode that uses the p2p.Service
func NewP2PNetwork(log logging.Logger, cfg config.Local, datadir string, phonebookAddresses []string, genesisID string, networkID protocol.NetworkID) (*P2PNetwork, error) {
	const readBufferLen = 2048

	// create Peerstore and add phonebook addresses
	addrInfo, malformedAddrs := peerstore.PeerInfoFromAddrs(phonebookAddresses)
	for malAddr, malErr := range malformedAddrs {
		log.Infof("Ignoring malformed phonebook address %s: %s", malAddr, malErr)
	}
	var pstore p2ppeerstore.Peerstore
	var err error
	if datadir == "" { // use ephemeral peerstore for testing
		pstore, err = pstoremem.NewPeerstore()
		if err != nil {
			return nil, err
		}
		for i := range addrInfo {
			pstore.AddAddrs(addrInfo[i].ID, addrInfo[i].Addrs, p2ppeerstore.AddressTTL)
		}
	} else {
		pstore, err = peerstore.NewPeerStore(context.Background(), filepath.Join(datadir, peerStorePath), addrInfo)
		if err != nil {
			return nil, err
		}
	}

	net := &P2PNetwork{
		log:          log,
		config:       cfg,
		genesisID:    genesisID,
		networkID:    networkID,
		topicTags:    map[protocol.Tag]string{"TX": p2p.TXTopicName},
		wsReadBuffer: make(chan IncomingMessage, readBufferLen),
		peers:        make(map[peer.ID]*wsPeer),
		peerStats:    make(map[peer.ID]*p2pPeerStats),
	}
	net.ctx, net.ctxCancel = context.WithCancel(context.Background())
	net.handlers.log = log
	net.broadcaster = broadcaster{
		ctx:                    net.ctx,
		log:                    log,
		config:                 cfg,
		broadcastQueueHighPrio: make(chan broadcastRequest, outgoingMessagesBufferSize),
		broadcastQueueBulk:     make(chan broadcastRequest, 100),
	}

	net.service, err = p2p.MakeService(log, cfg, pstore, net.streamHandler)
	if err != nil {
		return nil, err
	}

	err = net.setup()

	if err != nil {
		return nil, err
	}

	return net, nil
}

func (n *P2PNetwork) setup() error {
	if n.broadcaster.slowWritingPeerMonitorInterval == 0 {
		n.broadcaster.slowWritingPeerMonitorInterval = slowWritingPeerMonitorInterval
	}
	n.handlers.log = n.log
	return nil
}

// Start threads, listen on sockets.
func (n *P2PNetwork) Start() {
	n.wg.Add(1)
	go n.txTopicHandleLoop(n.ctx)
	n.wg.Add(1)
	go n.broadcaster.broadcastThread(n.wg.Done, n)
	n.service.DialPeers(n.config.GossipFanout)

	n.wg.Add(1)
	go n.meshThread()
}

// Stop closes sockets and stop threads.
func (n *P2PNetwork) Stop() {
	n.ctxCancel()
	n.service.Close()
	n.wg.Wait()
}

// streamHandler is a callback that the p2p package calls when a new peer connects and establishes a stream
// on a given protocol.
func (n *P2PNetwork) streamHandler(ctx context.Context, peer peer.ID, stream network.Stream) {
	if stream.Protocol() != p2p.AlgorandWsProtocol {
		// right now it only supports the legacy websocket protocol
		n.log.Warnf("unknown protocol %s", stream.Protocol())
		return
	}

	// get address for peer ID
	addr := stream.Conn().RemoteMultiaddr().String()
	if addr == "" {
		log.Errorf("Could not get address for peer %s", peer)
	}
	wsp := &wsPeer{
		wsPeerCore: makeP2PPeerCore(context.TODO(), n, addr, n.GetRoundTripper(), addr),
		conn:       &wsPeerConnP2PImpl{stream: stream},
	}
	wsp.init(n.config, outgoingMessagesBufferSize)
	n.peersLock.Lock()
	n.peers[peer] = wsp
	n.peersLock.Unlock()
	atomic.AddInt32(&n.peersChangeCounter, 1)
}

// called from wsPeer to report that it has closed
func (n *P2PNetwork) wsPeerRemoteClose(peer *wsPeer, reason disconnectReason) {
	n.peersLock.Lock()
	remotePeerID := peer.conn.(*wsPeerConnP2PImpl).stream.Conn().RemotePeer()
	delete(n.peers, remotePeerID)
	n.peersLock.Unlock()
	atomic.AddInt32(&n.peersChangeCounter, 1)
}

func (n *P2PNetwork) peerSnapshot(dest []*wsPeer) ([]*wsPeer, int32) {
	n.peersLock.RLock()
	defer n.peersLock.RUnlock()
	// based on wn.peerSnapshot
	if cap(dest) >= len(n.peers) {
		toClear := dest[len(n.peers):cap(dest)]
		for i := range toClear {
			if toClear[i] == nil {
				break
			}
			toClear[i] = nil
		}
		dest = dest[:len(n.peers)]
	} else {
		dest = make([]*wsPeer, len(n.peers))
	}
	i := 0
	for _, p := range n.peers {
		dest[i] = p
		i++
	}
	return dest, atomic.LoadInt32(&n.peersChangeCounter)
}

func (n *P2PNetwork) checkSlowWritingPeers() {}
func (n *P2PNetwork) getPeersChangeCounter() *int32 {
	return &n.peersChangeCounter
}

type wsPeerConnP2PImpl struct {
	stream network.Stream
}

func (c *wsPeerConnP2PImpl) RemoteAddrString() string {
	return c.stream.Conn().RemoteMultiaddr().String()
}

func (c *wsPeerConnP2PImpl) NextReader() (int, io.Reader, error) {
	// XXX maybe we should use our websockets library?
	var lenbuf [4]byte
	_, err := io.ReadFull(c.stream, lenbuf[:])
	if err != nil {
		return 0, nil, err
	}
	msglen := binary.BigEndian.Uint32(lenbuf[:])
	if msglen > MaxMessageLength {
		return 0, nil, fmt.Errorf("message too long: %d", msglen)
	}
	// return io.Reader that only reads the next msglen bytes
	return websocket.BinaryMessage, io.LimitReader(c.stream, int64(msglen)), nil
}

func (c *wsPeerConnP2PImpl) WriteMessage(_ int, buf []byte) error {
	// XXX maybe we should use our websockets library?
	// write encoding of the length
	var lenbuf [4]byte
	binary.BigEndian.PutUint32(lenbuf[:], uint32(len(buf)))
	_, err := c.stream.Write(lenbuf[:])
	if err != nil {
		return err
	}

	// write message
	_, err = c.stream.Write(buf)
	return err
}

func (c *wsPeerConnP2PImpl) CloseWithMessage([]byte, time.Time) error {
	return c.stream.Close()
}

func (c *wsPeerConnP2PImpl) SetReadLimit(int64) {
}

func (c *wsPeerConnP2PImpl) CloseWithoutFlush() error {
	return c.stream.Close()
}

func (c *wsPeerConnP2PImpl) UnderlyingConn() net.Conn { return nil }

func (n *P2PNetwork) meshThread() {
	defer n.wg.Done()
	timer := time.NewTicker(meshThreadInterval)
	defer timer.Stop()
	for {
		select {
		case <-timer.C:
			n.service.DialPeers(n.config.GossipFanout)
		case <-n.ctx.Done():
			return
		}
	}
}

func (n *P2PNetwork) txTopicHandleLoop(ctx context.Context) {
	defer n.wg.Done()
	sub, err := n.service.Subscribe(p2p.TXTopicName, n.txTopicValidator)
	if err != nil {
		n.log.Errorf("Failed to subscribe to topic %s: %v", p2p.TXTopicName, err)
		return
	}

	for {
		// XXX check ctx? or wait for pubsub to return error on sub.Next()?

		msg, err := sub.Next(ctx)
		if err != nil {
			if err != pubsub.ErrSubscriptionCancelled {
				n.log.Errorf("Error reading from subscription: %v, peerId:%s", err, n.service.Host().ID())
			}
			sub.Cancel()
			return
		}

		n.log.Infof("\nTxHandleLoop:\nReceived message from %s\nCurrent ID: %s\n\n", msg.ReceivedFrom, n.service.Host().ID())

		// handle TX message
		// from gossipsub's point of view, it's just waiting to hear back from the validator,
		// and txHandler does all its work in the validator, so we don't need to do anything here
		_ = msg
	}
}

// txTopicValidator calls txHandler to validate the TX message
func (n *P2PNetwork) txTopicValidator(ctx context.Context, peerID peer.ID, msg *pubsub.Message) pubsub.ValidationResult {
	inmsg := IncomingMessage{
		Sender:   msg.ReceivedFrom,
		Tag:      protocol.TxnTag,
		Data:     msg.Data,
		Net:      n,
		Received: time.Now().UnixNano(),
	}

	// if we sent the message, don't validate it
	if inmsg.Sender == n.service.Host().ID() {
		return pubsub.ValidationAccept
	}

	n.peerStatsMu.Lock() // XX probably don't wanna handle this here but adding it now just for testing the basic networking
	peerStats, ok := n.peerStats[peerID]
	if !ok {
		n.peerStats[peerID] = &p2pPeerStats{txReceived: 1}
	} else {
		peerStats.txReceived++
	}
	n.peerStatsMu.Unlock()

	outmsg := n.handlers.Handle(inmsg)
	// there was a decision made in the handler about this message
	switch outmsg.Action {
	case Ignore:
		return pubsub.ValidationIgnore
	case Disconnect:
		return pubsub.ValidationReject
	case Broadcast: // TxHandler.processIncomingTxn does not currently return this Action
		return pubsub.ValidationAccept
	default:
		n.log.Warnf("handler returned invalid action %d", outmsg.Action)
		return pubsub.ValidationIgnore
	}
}

// GetGenesisID implements GossipNode
func (n *P2PNetwork) GetGenesisID() string {
	return n.genesisID
}

// Address returns a string and whether that is a 'final' address or guessed.
func (n *P2PNetwork) Address() (string, bool) {
	addrs := n.service.Host().Addrs()
	if len(addrs) == 0 {
		return "", false
	}
	if len(addrs) > 1 {
		n.log.Infof("Multiple addresses found, using first one from %v", addrs)
	}
	return addrs[0].String(), true
}

// Broadcast sends a message.
// If except is not nil then we will not send it to that neighboring Peer.
// if wait is true then the call blocks until the packet has actually been sent to all neighbors.
func (n *P2PNetwork) Broadcast(ctx context.Context, tag protocol.Tag, data []byte, wait bool, except Peer) error {
	if _, ok := n.topicTags[tag]; ok {
		return n.service.Publish(context.TODO(), p2p.TXTopicName, data)
	}

	// use legacy network
	return n.broadcaster.BroadcastArray(ctx, []protocol.Tag{tag}, [][]byte{data}, wait, except)
}

// Relay message
func (n *P2PNetwork) Relay(ctx context.Context, tag protocol.Tag, data []byte, wait bool, except Peer) error {
	// handle TX messages with no-op: we assume GossipSub is already working to distribute valid TX messages
	if tag == protocol.TxnTag {
		return nil
	}

	// use legacy network
	return n.broadcaster.BroadcastArray(ctx, []protocol.Tag{tag}, [][]byte{data}, wait, except)
}

// Disconnect from a peer, probably due to protocol errors.
func (n *P2PNetwork) Disconnect(badnode Peer) {
	switch node := badnode.(type) {
	case peer.ID:
		err := n.service.Host().Network().ClosePeer(node)
		if err != nil {
			n.log.Warnf("Error disconnecting from peer %s: %v", node, err)
		}
	default:
		n.log.Warnf("Unknown peer type %T", badnode)
	}
}

// DisconnectPeers is used by testing
func (n *P2PNetwork) DisconnectPeers() {}

// RegisterHTTPHandler path accepts gorilla/mux path annotations
func (n *P2PNetwork) RegisterHTTPHandler(path string, handler http.Handler) {}

// RequestConnectOutgoing asks the system to actually connect to peers.
// `replace` optionally drops existing connections before making new ones.
// `quit` chan allows cancellation. TODO: use `context`
func (n *P2PNetwork) RequestConnectOutgoing(replace bool, quit <-chan struct{}) {
	// XXX catchup calls this
}

// GetPeers returns a list of Peers we could potentially send a direct message to.
func (n *P2PNetwork) GetPeers(options ...PeerOption) []Peer { return nil }

// RegisterHandlers adds to the set of given message handlers.
func (n *P2PNetwork) RegisterHandlers(dispatch []TaggedMessageHandler) {
	n.handlers.RegisterHandlers(dispatch)
}

// ClearHandlers deregisters all the existing message handlers.
func (n *P2PNetwork) ClearHandlers() {
	// XXX WebsocketNetwork actually just clears 3 handlers when this is called
	//n.handlers.ClearHandlers([]Tag{protocol.PingTag, protocol.PingReplyTag, protocol.NetPrioResponseTag})
}

// GetRoundTripper returns a Transport that would limit the number of outgoing connections.
func (n *P2PNetwork) GetRoundTripper() http.RoundTripper {
	return http.DefaultTransport
}

// OnNetworkAdvance notifies the network library that the agreement protocol was able to make a notable progress.
// this is the only indication that we have that we haven't formed a clique, where all incoming messages
// arrive very quickly, but might be missing some votes. The usage of this call is expected to have similar
// characteristics as with a watchdog timer.
func (n *P2PNetwork) OnNetworkAdvance() {}

// GetHTTPRequestConnection returns the underlying connection for the given request. Note that the request must be the same
// request that was provided to the http handler ( or provide a fallback Context() to that )
func (n *P2PNetwork) GetHTTPRequestConnection(request *http.Request) (conn net.Conn) { return nil }

// RegisterMessageInterest notifies the network library that this node
// wants to receive messages with the specified tag.  This will cause
// this node to send corresponding MsgOfInterest notifications to any
// newly connecting peers.  This should be called before the network
// is started.
func (n *P2PNetwork) RegisterMessageInterest(protocol.Tag) {}

// SubstituteGenesisID substitutes the "{genesisID}" with their network-specific genesisID.
func (n *P2PNetwork) SubstituteGenesisID(rawURL string) string {
	return strings.Replace(rawURL, "{genesisID}", n.genesisID, -1)
}
