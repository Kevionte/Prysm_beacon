package operations

import (
	"testing"

	"github.com/Kevionte/prysm_beacon/v2/testing/spectest/shared/capella/operations"
)

func TestMainnet_Capella_Operations_Withdrawals(t *testing.T) {
	operations.RunWithdrawalsTest(t, "mainnet")
}
