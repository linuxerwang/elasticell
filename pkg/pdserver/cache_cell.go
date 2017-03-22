// Copyright 2016 DeepFabric, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package pdserver

import (
	"github.com/deepfabric/elasticell/pkg/meta"
)

type cellRuntime struct {
	cell   *meta.CellMeta
	leader *meta.CellMeta
	// downPeers    []*pdpb.PeerStats
	pendingPeers []*meta.PeerMeta
}

func newCellRuntime(cell *meta.CellMeta) *cellRuntime {
	return &cellRuntime{
		cell: cell,
	}
}

func (cc *cellRuntime) getPendingPeer(peerID uint64) *meta.PeerMeta {
	for _, peer := range cc.pendingPeers {
		if peer.ID == peerID {
			return peer
		}
	}
	return nil
}

func (cc *cellRuntime) getPeer(peerID uint64) *meta.PeerMeta {
	for _, peer := range cc.cell.Peers {
		if peer.ID == peerID {
			return peer
		}
	}
	return nil
}

func (cc *cellRuntime) getID() uint64 {
	if cc.cell == nil {
		return 0
	}

	return cc.cell.ID
}

func (cc *cellRuntime) getPeers() []*meta.PeerMeta {
	if cc.cell == nil {
		return nil
	}

	return cc.cell.Peers
}

func (cc *cellRuntime) getStoreIDs() map[uint64]struct{} {
	peers := cc.getPeers()
	stores := make(map[uint64]struct{}, len(peers))
	for _, peer := range peers {
		stores[peer.StoreID] = struct{}{}
	}
	return stores
}