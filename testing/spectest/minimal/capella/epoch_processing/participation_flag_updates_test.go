package epoch_processing

import (
	"testing"

	"github.com/Kevionte/prysm_beacon/v2/testing/spectest/shared/capella/epoch_processing"
)

func TestMinimal_Capella_EpochProcessing_ParticipationFlag(t *testing.T) {
	epoch_processing.RunParticipationFlagUpdatesTests(t, "minimal")
}
