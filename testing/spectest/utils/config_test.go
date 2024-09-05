package utils

import (
	"testing"

	"github.com/Kevionte/prysm_beacon/v5/config/params"
	"github.com/Kevionte/prysm_beacon/v5/consensus-types/primitives"
	"github.com/Kevionte/prysm_beacon/v5/testing/require"
)

func TestConfig(t *testing.T) {
	require.NoError(t, SetConfig(t, "minimal"))
	require.Equal(t, primitives.Slot(8), params.BeaconConfig().SlotsPerEpoch)
	require.NoError(t, SetConfig(t, "mainnet"))
	require.Equal(t, primitives.Slot(32), params.BeaconConfig().SlotsPerEpoch)
}
