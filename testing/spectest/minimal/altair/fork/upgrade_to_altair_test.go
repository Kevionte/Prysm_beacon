package fork

import (
	"testing"

	"github.com/Kevionte/prysm_beacon/v5/testing/spectest/shared/altair/fork"
)

func TestMinimal_Altair_UpgradeToAltair(t *testing.T) {
	fork.RunUpgradeToAltair(t, "minimal")
}
