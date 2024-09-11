package p2p

import (
	"time"

	"github.com/Kevionte/prysm_beacon/v1beacon-chain/forkchoice"
	"github.com/Kevionte/prysm_beacon/v1consensus-types/primitives"
	ethpb "github.com/Kevionte/prysm_beacon/v1proto/prysm/v1alpha1"
	"github.com/Kevionte/prysm_beacon/v1time/slots"
)

type mockChain struct {
	currentFork     *ethpb.Fork
	genesisValsRoot [32]byte
	genesisTime     time.Time
}

func (m *mockChain) ForkChoicer() forkchoice.ForkChoicer {
	return nil
}

func (m *mockChain) CurrentFork() *ethpb.Fork {
	return m.currentFork
}

func (m *mockChain) GenesisValidatorsRoot() [32]byte {
	return m.genesisValsRoot
}

func (m *mockChain) GenesisTime() time.Time {
	return m.genesisTime
}

func (m *mockChain) CurrentSlot() primitives.Slot {
	return slots.SinceGenesis(m.genesisTime)
}
