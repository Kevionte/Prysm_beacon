package cache

import (
	"errors"
	"sync"

	forkchoicetypes "github.com/Kevionte/prysm_beacon/v2/beacon-chain/forkchoice/types"
	"github.com/Kevionte/prysm_beacon/v2/consensus-types/primitives"
)

type AttestationConsensusData struct {
	Slot     primitives.Slot
	HeadRoot []byte
	Target   forkchoicetypes.Checkpoint
	Source   forkchoicetypes.Checkpoint
}

// AttestationCache stores cached results of AttestationData requests.
type AttestationCache struct {
	a *AttestationConsensusData
	sync.RWMutex
}

// NewAttestationCache creates a new instance of AttestationCache.
func NewAttestationCache() *AttestationCache {
	return &AttestationCache{}
}

// Get retrieves cached attestation data, recording a cache hit or miss. This method is lock free.
func (c *AttestationCache) Get() *AttestationConsensusData {
	return c.a
}

// Put adds a response to the cache. This method is lock free.
func (c *AttestationCache) Put(a *AttestationConsensusData) error {
	if a == nil {
		return errors.New("attestation cannot be nil")
	}
	c.a = a
	return nil
}
