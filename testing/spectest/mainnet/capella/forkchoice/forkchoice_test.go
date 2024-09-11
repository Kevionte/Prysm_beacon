package forkchoice

import (
	"testing"

	"github.com/Kevionte/prysm_beacon/v2/runtime/version"
	"github.com/Kevionte/prysm_beacon/v2/testing/spectest/shared/common/forkchoice"
)

func TestMainnet_Capella_Forkchoice(t *testing.T) {
	forkchoice.Run(t, "mainnet", version.Capella)
}
