package testutil

import (
	"context"

	"github.com/Kevionte/prysm_beacon/v1encoding/bytesutil"

	"github.com/Kevionte/prysm_beacon/v1beacon-chain/state"
	"github.com/Kevionte/prysm_beacon/v1consensus-types/primitives"
)

// MockStater is a fake implementation of lookup.Stater.
type MockStater struct {
	BeaconState       state.BeaconState
	StateProviderFunc func(ctx context.Context, stateId []byte) (state.BeaconState, error)
	BeaconStateRoot   []byte
	StatesBySlot      map[primitives.Slot]state.BeaconState
	StatesByRoot      map[[32]byte]state.BeaconState
}

// State --
func (m *MockStater) State(ctx context.Context, id []byte) (state.BeaconState, error) {
	if m.StateProviderFunc != nil {
		return m.StateProviderFunc(ctx, id)
	}

	if m.BeaconState != nil {
		return m.BeaconState, nil
	}

	return m.StatesByRoot[bytesutil.ToBytes32(id)], nil
}

// StateRoot --
func (m *MockStater) StateRoot(context.Context, []byte) ([]byte, error) {
	return m.BeaconStateRoot, nil
}

// StateBySlot --
func (m *MockStater) StateBySlot(_ context.Context, s primitives.Slot) (state.BeaconState, error) {
	return m.StatesBySlot[s], nil
}
