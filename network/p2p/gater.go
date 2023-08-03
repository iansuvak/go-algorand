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
	"github.com/libp2p/go-libp2p/core/control"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	ma "github.com/multiformats/go-multiaddr"
)

// methods defined here implement libp2p's ConnectionGater interface

// InterceptPeerDial implements connmgr.ConnectionGater
func (s *Service) InterceptPeerDial(p peer.ID) (allow bool) {
	return s.host.ID() != p
}

// InterceptAddrDial implements connmgr.ConnectionGater
// Trivially allow all outgoing connections right now
// TODO: potentially implement preventing re-dialing known bad peers and prevent dialing if we are already at max peers
func (_ *Service) InterceptAddrDial(peer.ID, ma.Multiaddr) (allow bool) {
	return true

}

// InterceptAccept tests whether an incipient inbound connection is allowed.
func (_ *Service) InterceptAccept(network.ConnMultiaddrs) (allow bool) {
	return true
}

// InterceptSecured tests whether a given connection, now authenticated,
// is allowed.
//
// This is called by the upgrader, after it has performed the security
// handshake, and before it negotiates the muxer, or by the directly by the
// transport, at the exact same checkpoint.
func (_ *Service) InterceptSecured(network.Direction, peer.ID, network.ConnMultiaddrs) (allow bool) {
	return true
}

// InterceptUpgraded tests whether a fully capable connection is allowed.
//
// At this point, the connection a multiplexer has been selected.
// When rejecting a connection, the gater can return a DisconnectReason.
// Refer to the godoc on the ConnectionGater type for more information.
func (_ *Service) InterceptUpgraded(network.Conn) (allow bool, reason control.DisconnectReason) {
	return true, control.DisconnectReason(0)
}
