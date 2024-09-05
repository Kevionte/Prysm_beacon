package forkchoice

import (
	"testing"

	"github.com/Kevionte/prysm_beacon/v5/runtime/version"
	"github.com/Kevionte/prysm_beacon/v5/testing/spectest/shared/common/forkchoice"
)

func TestMainnet_Deneb_Forkchoice(t *testing.T) {
	forkchoice.Run(t, "mainnet", version.Deneb)
}
