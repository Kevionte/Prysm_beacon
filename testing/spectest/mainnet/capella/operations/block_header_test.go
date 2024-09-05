package operations

import (
	"testing"

	"github.com/Kevionte/prysm_beacon/v5/testing/spectest/shared/capella/operations"
)

func TestMainnet_Capella_Operations_BlockHeader(t *testing.T) {
	operations.RunBlockHeaderTest(t, "mainnet")
}
