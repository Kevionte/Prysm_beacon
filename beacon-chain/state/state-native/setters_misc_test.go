package state_native_test

import (
	"testing"

	state_native "github.com/Kevionte/prysm_beacon/v1beacon-chain/state/state-native"
	"github.com/Kevionte/prysm_beacon/v1config/params"
	"github.com/Kevionte/prysm_beacon/v1encoding/bytesutil"
	ethpb "github.com/Kevionte/prysm_beacon/v1proto/prysm/v1alpha1"
	"github.com/Kevionte/prysm_beacon/v1testing/require"
)

func BenchmarkAppendHistoricalRoots(b *testing.B) {
	st, err := state_native.InitializeFromProtoPhase0(&ethpb.BeaconState{})
	require.NoError(b, err)

	max := params.BeaconConfig().HistoricalRootsLimit
	if max < 2 {
		b.Fatalf("HistoricalRootsLimit is less than 2: %d", max)
	}

	root := bytesutil.ToBytes32([]byte{0, 1, 2, 3, 4, 5})
	for i := uint64(0); i < max-2; i++ {
		err := st.AppendHistoricalRoots(root)
		require.NoError(b, err)
	}

	ref := st.Copy()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := ref.AppendHistoricalRoots(root)
		require.NoError(b, err)
		ref = st.Copy()
	}
}

func BenchmarkAppendHistoricalSummaries(b *testing.B) {
	st, err := state_native.InitializeFromProtoCapella(&ethpb.BeaconStateCapella{})
	require.NoError(b, err)

	max := params.BeaconConfig().HistoricalRootsLimit
	if max < 2 {
		b.Fatalf("HistoricalRootsLimit is less than 2: %d", max)
	}

	for i := uint64(0); i < max-2; i++ {
		err := st.AppendHistoricalSummaries(&ethpb.HistoricalSummary{})
		require.NoError(b, err)
	}

	ref := st.Copy()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := ref.AppendHistoricalSummaries(&ethpb.HistoricalSummary{})
		require.NoError(b, err)
		ref = st.Copy()
	}
}
