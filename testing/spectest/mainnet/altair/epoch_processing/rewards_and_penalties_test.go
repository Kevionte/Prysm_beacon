package epoch_processing

import (
	"testing"

	"github.com/Kevionte/prysm_beacon/v2/testing/spectest/shared/altair/epoch_processing"
)

func TestMainnet_Altair_EpochProcessing_RewardsAndPenalties(t *testing.T) {
	epoch_processing.RunRewardsAndPenaltiesTests(t, "mainnet")
}
