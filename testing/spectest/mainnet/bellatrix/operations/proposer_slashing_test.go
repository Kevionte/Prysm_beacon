package operations

import (
	"testing"

	"github.com/Kevionte/prysm_beacon/v1testing/spectest/shared/bellatrix/operations"
)

func TestMainnet_Bellatrix_Operations_ProposerSlashing(t *testing.T) {
	operations.RunProposerSlashingTest(t, "mainnet")
}
