package operations

import (
	"testing"

	"github.com/Kevionte/prysm_beacon/v5/testing/spectest/shared/deneb/operations"
)

func TestMainnet_Deneb_Operations_Deposit(t *testing.T) {
	operations.RunDepositTest(t, "mainnet")
}
