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
	"crypto/rand"
	"fmt"
	"runtime"

	"github.com/algorand/go-algorand/config"
	"github.com/algorand/go-algorand/logging"
	"github.com/algorand/go-deadlock"
	"github.com/multiformats/go-multiaddr"

	"github.com/libp2p/go-libp2p"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/host"
	peerstore "github.com/libp2p/go-libp2p/core/peer"

	"github.com/libp2p/go-libp2p/p2p/muxer/yamux"
	"github.com/libp2p/go-libp2p/p2p/transport/tcp"
)

// Service manages integration with libp2p
type Service struct {
	log              logging.Logger
	host             host.Host
	privKey          crypto.PrivKey
	streams          *streamManager
	pubsub           *pubsub.PubSub
	pubsubCtx        context.Context
	initialPhonebook []multiaddr.Multiaddr

	topics   map[string]*pubsub.Topic
	topicsMu deadlock.Mutex
}

// AlgorandWsProtocol defines a libp2p protocol name for algorand's websockets messages
const AlgorandWsProtocol = "/algorand-ws/1.0.0"

// MakeService creates a P2P service instance
func MakeService(log logging.Logger, cfg config.Local, phonebookAddresses []string, streamHandler StreamHandler) (*Service, error) {
	// ephemeral key for peer ID
	privKey, _, err := crypto.GenerateEd25519Key(rand.Reader)
	if err != nil {
		return nil, err
	}

	// muxer supports tweaking fields from yamux.Config
	ymx := *yamux.DefaultTransport
	// user-agent copied from wsNetwork.go
	version := config.GetCurrentVersion()
	ua := fmt.Sprintf("algod/%d.%d (%s; commit=%s; %d) %s(%s)", version.Major, version.Minor, version.Channel, version.CommitHash, version.BuildNumber, runtime.GOOS, runtime.GOARCH)

	h, err := libp2p.New(
		libp2p.Identity(privKey),
		libp2p.UserAgent(ua),
		libp2p.Transport(tcp.NewTCPTransport),
		libp2p.Muxer("/yamux/1.0.0", &ymx),
		// libp2p.PeerStore(XXX), // default is empty in-memory peerstore
		libp2p.ListenAddrStrings("/ip4/0.0.0.0/tcp/0"), // need to include this to be able to get address for self
		// libp2p.ConnectionGater(XXX),
	)
	if err != nil {
		return nil, err
	}

	sm := makeStreamManager(h, streamHandler)
	h.SetStreamHandler(AlgorandWsProtocol, sm.streamHandler)

	psCtx := context.TODO()
	ps, err := makePubSub(psCtx, cfg, h)
	if err != nil {
		return nil, err
	}

	multiAdrPhoneBook := make([]multiaddr.Multiaddr, 0)

	for _, addr := range phonebookAddresses {
		multiAdr, err := multiaddr.NewMultiaddr(addr)
		if err != nil {
			return nil, err
		}
		multiAdrPhoneBook = append(multiAdrPhoneBook, multiAdr)
	}

	return &Service{
		log:              log,
		privKey:          privKey,
		host:             h,
		streams:          sm,
		pubsub:           ps,
		pubsubCtx:        psCtx,
		topics:           make(map[string]*pubsub.Topic),
		initialPhonebook: multiAdrPhoneBook,
	}, nil
}

// Close shuts down the P2P service
func (s *Service) Close() error {
	return s.host.Close()
}

// Host returns the libp2p host
func (s *Service) Host() host.Host {
	return s.host
}

// MakeInitialConnections attempts to establish connections to the provided phonebook addresses
func (s *Service) MakeInitialConnections() {
	for _, addr := range s.initialPhonebook {
		addrInfo, err := peerstore.AddrInfoFromP2pAddr(addr)
		if err != nil {
			s.log.Warnf("Failed to get address info for phonebook address: %s, err: %s", addr, err)
			continue
		}
		err = s.Host().Connect(context.TODO(), *addrInfo)
		if err != nil {
			s.log.Warnf("Failed to connect to phonebook address: %s, err: %s", addr, err)
		}
	}
}
