package operations

import (
	"testing"

	"github.com/Kevionte/prysm_beacon/v1testing/spectest/shared/altair/operations"
)

func TestMinimal_Altair_Operations_VoluntaryExit(t *testing.T) {
	operations.RunVoluntaryExitTest(t, "minimal")
}
