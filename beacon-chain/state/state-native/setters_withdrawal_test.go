package state_native

import (
	"testing"

	"github.com/Kevionte/prysm_beacon/v1beacon-chain/state/state-native/types"
	"github.com/Kevionte/prysm_beacon/v1consensus-types/primitives"
	"github.com/Kevionte/prysm_beacon/v1runtime/version"
	"github.com/Kevionte/prysm_beacon/v1testing/require"
)

func TestSetNextWithdrawalIndex(t *testing.T) {
	s := BeaconState{
		version:             version.Capella,
		nextWithdrawalIndex: 3,
		dirtyFields:         make(map[types.FieldIndex]bool),
	}
	require.NoError(t, s.SetNextWithdrawalIndex(5))
	require.Equal(t, uint64(5), s.nextWithdrawalIndex)
	require.Equal(t, true, s.dirtyFields[types.NextWithdrawalIndex])
}

func TestSetNextWithdrawalValidatorIndex(t *testing.T) {
	s := BeaconState{
		version:                      version.Capella,
		nextWithdrawalValidatorIndex: 3,
		dirtyFields:                  make(map[types.FieldIndex]bool),
	}
	require.NoError(t, s.SetNextWithdrawalValidatorIndex(5))
	require.Equal(t, primitives.ValidatorIndex(5), s.nextWithdrawalValidatorIndex)
	require.Equal(t, true, s.dirtyFields[types.NextWithdrawalValidatorIndex])
}

func TestSetNextWithdrawalIndex_Deneb(t *testing.T) {
	s := BeaconState{
		version:             version.Deneb,
		nextWithdrawalIndex: 3,
		dirtyFields:         make(map[types.FieldIndex]bool),
	}
	require.NoError(t, s.SetNextWithdrawalIndex(5))
	require.Equal(t, uint64(5), s.nextWithdrawalIndex)
	require.Equal(t, true, s.dirtyFields[types.NextWithdrawalIndex])
}

func TestSetNextWithdrawalValidatorIndex_Deneb(t *testing.T) {
	s := BeaconState{
		version:                      version.Deneb,
		nextWithdrawalValidatorIndex: 3,
		dirtyFields:                  make(map[types.FieldIndex]bool),
	}
	require.NoError(t, s.SetNextWithdrawalValidatorIndex(5))
	require.Equal(t, primitives.ValidatorIndex(5), s.nextWithdrawalValidatorIndex)
	require.Equal(t, true, s.dirtyFields[types.NextWithdrawalValidatorIndex])
}
