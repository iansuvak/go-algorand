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
	"fmt"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/algorand/go-algorand/config"
	"github.com/algorand/go-algorand/logging"
	"github.com/algorand/go-algorand/network/p2p"
	"github.com/algorand/go-algorand/protocol"

	"github.com/labstack/gommon/log"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	p2pprotocol "github.com/libp2p/go-libp2p/core/protocol"
)

// P2PNetwork implements the GossipNode interface
type P2PNetwork struct {
	service   *p2p.Service
	log       logging.Logger
	genesisID string
	networkID protocol.NetworkID

	wg sync.WaitGroup

	// which tags to use with libp2p's GossipSub, mapped to topic names
	topicTags map[protocol.Tag]string

	handlers Multiplexer
}

// NewP2PNetwork returns an instance of GossipNode that uses the p2p.Service
func NewP2PNetwork(log logging.Logger, cfg config.Local, phonebookAddresses []string, genesisID string, networkID protocol.NetworkID) (*P2PNetwork, error) {
	p2pService, err := p2p.MakeService(log, cfg, phonebookAddresses)
	if err != nil {
		return nil, err
	}
	net := &P2PNetwork{
		service:   p2pService,
		log:       log,
		genesisID: genesisID,
		networkID: networkID,
		topicTags: map[protocol.Tag]string{"TX": p2p.TXTopicName},
	}
	net.handlers.log = log

	return net, nil
}

// Start threads, listen on sockets.
func (n *P2PNetwork) Start() {

	n.wg.Add(1)
	go n.txTopicHandleLoop(context.TODO())
}

// Close sockets. Stop threads.
func (n *P2PNetwork) Stop() {
	n.service.Close()
	n.wg.Wait()
}

// streamHandler is a callback that the p2p package calls when a new peer connects and establishes a stream
// on a given protocol.
func (n *P2PNetwork) streamHandler(ctx context.Context, proto p2pprotocol.ID, peer peer.ID, stream network.Stream) error {
	if proto != p2p.AlgorandWsProtocol {
		// right now it only supports the legacy websocket protocol
		return fmt.Errorf("unknown protocol %s", proto)
	}

	n.wg.Add(1)
	go func() {
		defer n.wg.Done()

	}()
	return nil
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
				log.Errorf("Error reading from subscription: %v", err)
			}
			sub.Cancel()
			return
		}

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
	outmsg := n.handlers.Handle(inmsg)
	if !outmsg.ValidationQueued {
		switch outmsg.Action {
		case Ignore:
			return pubsub.ValidationIgnore
		case Disconnect:
			return pubsub.ValidationReject
		case Broadcast:
			return pubsub.ValidationAccept
		}
	}
	// txHandler queued txn for signature & pool validation
	// XXX incorporate feedback for txHandler to tell us when a queued message is approved or rejected
	return pubsub.ValidationAccept
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
	return n.legacyBroadcast(ctx, tag, data, wait, except)
}

func (n *P2PNetwork) legacyBroadcast(ctx context.Context, tag protocol.Tag, data []byte, wait bool, except Peer) error {
	return nil
}

// Relay message
func (n *P2PNetwork) Relay(ctx context.Context, tag protocol.Tag, data []byte, wait bool, except Peer) error {
	// handle TX messages with no-op: we assume GossipSub is already working to distribute valid TX messages
	if tag == protocol.TxnTag {
		return nil
	}

	// use legacy network
	return n.legacyBroadcast(ctx, tag, data, wait, except)
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

// RegisterHTTPHandler path accepts gorilla/mux path annotations
func (n *P2PNetwork) RegisterHTTPHandler(path string, handler http.Handler) {}

// RequestConnectOutgoing asks the system to actually connect to peers.
// `replace` optionally drops existing connections before making new ones.
// `quit` chan allows cancellation. TODO: use `context`
func (n *P2PNetwork) RequestConnectOutgoing(replace bool, quit <-chan struct{}) {
	// XXX catchup calls this
}

// Get a list of Peers we could potentially send a direct message to.
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
