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

package p2p

import (
	"context"
	"io"
	"log"
	"sync"

	"github.com/algorand/go-deadlock"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/protocol"
	"github.com/multiformats/go-multiaddr"
)

// streamManager implements network.Notifiee to create and manage streams for use with non-gossipsub protocols.
// XXX could also manage HTTP streams for e.g. fast catchup
type streamManager struct {
	host    host.Host
	handler StreamHandler

	streams     map[peer.ID]network.Stream
	streamsLock deadlock.Mutex
}

// StreamHandler is called when a new bidirectional stream for a given protocol and peer is opened.
type StreamHandler func(context.Context, peer.ID, network.Stream)

func makeStreamManager(h host.Host, handler StreamHandler) *streamManager {
	return &streamManager{
		host:    h,
		handler: handler,
		streams: make(map[peer.ID]network.Stream),
	}
}

// streamHandler is called by libp2p when a new stream is accepted
func (n *streamManager) streamHandler(stream network.Stream) {
	n.streamsLock.Lock()
	defer n.streamsLock.Unlock()

	// could use stream.ID() for tracking; unique across all conns and peers

	remotePeer := stream.Conn().RemotePeer()
	if oldStream, ok := n.streams[remotePeer]; ok {
		// there's already a stream, for some reason, check if it's still open
		buf := []byte{} // empty buffer for checking
		_, err := oldStream.Read(buf)
		if err != nil {
			if err == io.EOF {
				// old stream was closed by the peer
				log.Printf("Old stream with %s was closed", remotePeer)
			} else {
				// an error occurred while checking the old stream
				log.Printf("Failed to check old stream with %s: %v", remotePeer, err)
			}
			n.streams[stream.Conn().RemotePeer()] = stream
			n.handler(context.TODO(), remotePeer, stream)
			return
		}
		// otherwise, the old stream is still open, so we can close the new one
		stream.Close()
		return
	}
	// no old stream
	n.streams[stream.Conn().RemotePeer()] = stream
	n.handler(context.TODO(), remotePeer, stream)
}

// Connected is called when a connection is opened
func (n *streamManager) Connected(net network.Network, conn network.Conn) {
	remotePeer := conn.RemotePeer()
	localPeer := n.host.ID()

	// ensure that only one of the peers initiates the stream
	if localPeer > remotePeer {
		return
	}

	n.streamsLock.Lock()
	defer n.streamsLock.Unlock()
	_, ok := n.streams[remotePeer]
	if ok {
		return // there's already an active stream with this peer for our protocol
	}

	stream, err := n.host.NewStream(context.Background(), conn.RemotePeer(), protocol.ID(AlgorandWsProtocol))
	if err != nil {
		log.Printf("Failed to open stream to %s: %v", conn.RemotePeer(), err)
		return
	}
	n.streams[conn.RemotePeer()] = stream
}

// Disconnected is called when a connection is closed
func (n *streamManager) Disconnected(net network.Network, conn network.Conn) {
	n.streamsLock.Lock()
	defer n.streamsLock.Unlock()

	stream, ok := n.streams[conn.RemotePeer()]
	if ok {
		stream.Close()
		delete(n.streams, conn.RemotePeer())
	}
}

// Listen is called when network starts listening on an addr
func (n *streamManager) Listen(net network.Network, addr multiaddr.Multiaddr) {}

// ListenClose is called when network stops listening on an addr
func (n *streamManager) ListenClose(net network.Network, addr multiaddr.Multiaddr) {}
