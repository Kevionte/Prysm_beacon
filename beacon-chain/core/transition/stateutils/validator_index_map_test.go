package stateutils_test

import (
	"testing"

	"github.com/Kevionte/prysm_beacon/v2/beacon-chain/core/transition/stateutils"
	state_native "github.com/Kevionte/prysm_beacon/v2/beacon-chain/state/state-native"
	fieldparams "github.com/Kevionte/prysm_beacon/v2/config/fieldparams"
	"github.com/Kevionte/prysm_beacon/v2/consensus-types/primitives"
	"github.com/Kevionte/prysm_beacon/v2/encoding/bytesutil"
	ethpb "github.com/Kevionte/prysm_beacon/v2/proto/prysm/v1alpha1"
	"github.com/Kevionte/prysm_beacon/v2/testing/assert"
	"github.com/Kevionte/prysm_beacon/v2/testing/require"
)

func TestValidatorIndexMap_OK(t *testing.T) {
	base := &ethpb.BeaconState{
		Validators: []*ethpb.Validator{
			{
				PublicKey: []byte("zero"),
			},
			{
				PublicKey: []byte("one"),
			},
		},
	}
	state, err := state_native.InitializeFromProtoPhase0(base)
	require.NoError(t, err)

	tests := []struct {
		key [fieldparams.BLSPubkeyLength]byte
		val primitives.ValidatorIndex
		ok  bool
	}{
		{
			key: bytesutil.ToBytes48([]byte("zero")),
			val: 0,
			ok:  true,
		}, {
			key: bytesutil.ToBytes48([]byte("one")),
			val: 1,
			ok:  true,
		}, {
			key: bytesutil.ToBytes48([]byte("no")),
			val: 0,
			ok:  false,
		},
	}

	m := stateutils.ValidatorIndexMap(state.Validators())
	for _, tt := range tests {
		result, ok := m[tt.key]
		assert.Equal(t, tt.val, result)
		assert.Equal(t, tt.ok, ok)
	}
}
