package net

import (
	"math"
	"math/rand"
)

// ChainSyncPeersFilter will filter some peers randomly
type ChainSyncPeersFilter struct {
}

// Filter implemets PeerFilterAlgorithm interface
func (filter *ChainSyncPeersFilter) Filter(peers PeersSlice) PeersSlice {
	if len(peers) == 0 {
		return peers
	}
	selection := int(math.Sqrt(float64(len(peers))))
	return peers[:selection]
}

// RandomPeerFilter will filter a peer randomly
type RandomPeerFilter struct {
}

// Filter implemets PeerFilterAlgorithm interface
func (filter *RandomPeerFilter) Filter(peers PeersSlice) PeersSlice {
	if len(peers) == 0 {
		return peers
	}

	selection := rand.Intn(len(peers))
	return peers[selection : selection+1]
}
