package finality

import (
	"testing"

	"github.com/Kevionte/prysm_beacon/v5/testing/spectest/shared/deneb/finality"
)

func TestMainnet_Deneb_Finality(t *testing.T) {
	finality.RunFinalityTest(t, "mainnet")
}
