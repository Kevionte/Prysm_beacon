package operations

import (
	"testing"

	"github.com/Kevionte/prysm_beacon/v1testing/spectest/shared/bellatrix/operations"
)

func TestMinimal_Bellatrix_Operations_PayloadExecution(t *testing.T) {
	operations.RunExecutionPayloadTest(t, "minimal")
}
