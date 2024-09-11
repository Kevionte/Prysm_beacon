package operations

import (
	"testing"

	"github.com/Kevionte/prysm_beacon/v2/testing/spectest/shared/bellatrix/operations"
)

func TestMinimal_Bellatrix_Operations_Attestation(t *testing.T) {
	operations.RunAttestationTest(t, "minimal")
}
