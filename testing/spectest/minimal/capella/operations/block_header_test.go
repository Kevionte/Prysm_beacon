package operations

import (
	"testing"

	"github.com/Kevionte/prysm_beacon/v2/testing/spectest/shared/capella/operations"
)

func TestMinimal_Capella_Operations_BlockHeader(t *testing.T) {
	operations.RunBlockHeaderTest(t, "minimal")
}
