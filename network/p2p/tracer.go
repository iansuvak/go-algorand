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
	"github.com/algorand/go-algorand/logging"
	"github.com/algorand/go-algorand/logging/telemetryspec"
	"github.com/algorand/go-algorand/util/metrics"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/protocol"
)

// This tracer is used to implement metrics collection for messages received
// and broadcasted through gossipsub.
type gossipTracer struct {
	host host.Host
	log  logging.Logger
}

var p2pAddPeer = metrics.MakeCounter(metrics.MetricName{Name: "algod_p2p_add_peer", Description: "Number of peers added"})
var p2pRemovePeer = metrics.MakeCounter(metrics.MetricName{Name: "algod_p2p_remove_peer", Description: "Number of peers removed"})
var p2pGraftPeer = metrics.MakeCounter(metrics.MetricName{Name: "algod_p2p_graft_peer", Description: "Number of peers grafted"})
var p2pPrunePeer = metrics.MakeCounter(metrics.MetricName{Name: "algod_p2p_prune_peer", Description: "Number of peers pruned"})
var p2pValidateMessage = metrics.MakeCounter(metrics.MetricName{Name: "algod_p2p_validate_message", Description: "Number of messages validated"})
var p2pDeliverMessage = metrics.MakeCounter(metrics.MetricName{Name: "algod_p2p_deliver_message", Description: "Number of messages delivered"})
var p2pRejectMessage = metrics.MakeCounter(metrics.MetricName{Name: "algod_p2p_reject_message", Description: "Number of messages rejected"})
var p2pDuplicateMessage = metrics.MakeCounter(metrics.MetricName{Name: "algod_p2p_duplicate_message", Description: "Number of duplicate messages filtered"})

// AddPeer event
func (g gossipTracer) AddPeer(p peer.ID, proto protocol.ID) {
	p2pAddPeer.Inc(nil)
}

// RemovePeer event
func (g gossipTracer) RemovePeer(p peer.ID) {
	p2pRemovePeer.Inc(nil)
}

// Join event is unimplemented for now
func (g gossipTracer) Join(topic string) {
}

// Leave event is unimplemented for now
func (g gossipTracer) Leave(topic string) {
}

// Graft event
func (g gossipTracer) Graft(p peer.ID, topic string) {
	p2pGraftPeer.Inc(nil)
	details := telemetryspec.P2PPeerEventDetails{
		PeerID: p.String(),
		Topic:  topic,
	}
	g.log.EventWithDetails(telemetryspec.GossipSub, telemetryspec.GraftPeerEvent, details)
}

// Prune event
func (g gossipTracer) Prune(p peer.ID, topic string) {
	p2pPrunePeer.Inc(nil)
	details := telemetryspec.P2PPeerEventDetails{
		PeerID: p.String(),
		Topic:  topic,
	}
	g.log.EventWithDetails(telemetryspec.GossipSub, telemetryspec.PrunePeerEvent, details)
}

// ValidateMessage event
func (g gossipTracer) ValidateMessage(msg *pubsub.Message) {
	p2pValidateMessage.Inc(nil)
}

// DeliverMessage event
func (g gossipTracer) DeliverMessage(msg *pubsub.Message) {
	p2pDeliverMessage.Inc(nil)
}

// RejectMessage event
func (g gossipTracer) RejectMessage(msg *pubsub.Message, reason string) {
	p2pRejectMessage.Inc(nil)
}

// DuplicateMessage event is unimplemented for now
func (g gossipTracer) DuplicateMessage(msg *pubsub.Message) {
	p2pDuplicateMessage.Inc(nil)
}

// UndeliverableMessage event is unimplemented for now
func (g gossipTracer) UndeliverableMessage(msg *pubsub.Message) {
}

// ThrottlePeer event is unimplemented for now
func (g gossipTracer) ThrottlePeer(p peer.ID) {
}

// RecvRPC event is unimplemented for now
func (g gossipTracer) RecvRPC(rpc *pubsub.RPC) {
}

// SendRPC event is unimplemented for now
func (g gossipTracer) SendRPC(rpc *pubsub.RPC, p peer.ID) {
}

// DropRPC event is unimplemented for now
func (g gossipTracer) DropRPC(rpc *pubsub.RPC, p peer.ID) {
}
