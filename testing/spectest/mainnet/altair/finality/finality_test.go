package finality

import (
	"testing"

	"github.com/Kevionte/prysm_beacon/v5/testing/spectest/shared/altair/finality"
)

func TestMainnet_Altair_Finality(t *testing.T) {
	finality.RunFinalityTest(t, "mainnet")
}
