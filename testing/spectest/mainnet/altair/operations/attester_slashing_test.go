package operations

import (
	"testing"

	"github.com/Kevionte/prysm_beacon/v1testing/spectest/shared/altair/operations"
)

func TestMainnet_Altair_Operations_AttesterSlashing(t *testing.T) {
	operations.RunAttesterSlashingTest(t, "mainnet")
}
