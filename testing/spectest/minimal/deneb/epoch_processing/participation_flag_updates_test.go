package epoch_processing

import (
	"testing"

	"github.com/Kevionte/prysm_beacon/v1testing/spectest/shared/deneb/epoch_processing"
)

func TestMinimal_Deneb_EpochProcessing_ParticipationFlag(t *testing.T) {
	epoch_processing.RunParticipationFlagUpdatesTests(t, "minimal")
}
