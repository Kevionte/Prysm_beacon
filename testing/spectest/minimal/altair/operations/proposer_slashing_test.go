package operations

import (
	"testing"

	"github.com/Kevionte/prysm_beacon/v5/testing/spectest/shared/altair/operations"
)

func TestMinimal_Altair_Operations_ProposerSlashing(t *testing.T) {
	operations.RunProposerSlashingTest(t, "minimal")
}
