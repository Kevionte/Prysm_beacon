package epoch_processing

import (
	"testing"

	"github.com/Kevionte/prysm_beacon/v5/testing/spectest/shared/phase0/epoch_processing"
)

func TestMinimal_Phase0_EpochProcessing_Slashings(t *testing.T) {
	epoch_processing.RunSlashingsTests(t, "minimal")
}
