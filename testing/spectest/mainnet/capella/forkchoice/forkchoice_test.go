package forkchoice

import (
	"testing"

	"github.com/Kevionte/prysm_beacon/v1runtime/version"
	"github.com/Kevionte/prysm_beacon/v1testing/spectest/shared/common/forkchoice"
)

func TestMainnet_Capella_Forkchoice(t *testing.T) {
	forkchoice.Run(t, "mainnet", version.Capella)
}
