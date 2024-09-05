package rewards

import (
	"testing"

	"github.com/Kevionte/prysm_beacon/v5/testing/spectest/shared/altair/rewards"
)

func TestMainnet_Altair_Rewards(t *testing.T) {
	rewards.RunPrecomputeRewardsAndPenaltiesTests(t, "mainnet")
}
