package rewards

import (
	"testing"

	"github.com/Kevionte/prysm_beacon/v5/testing/spectest/shared/capella/rewards"
)

func TestMainnet_Capella_Rewards(t *testing.T) {
	rewards.RunPrecomputeRewardsAndPenaltiesTests(t, "mainnet")
}
