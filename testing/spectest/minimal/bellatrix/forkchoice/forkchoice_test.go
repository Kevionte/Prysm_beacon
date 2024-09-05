package forkchoice

import (
	"testing"

	"github.com/Kevionte/prysm_beacon/v5/runtime/version"
	"github.com/Kevionte/prysm_beacon/v5/testing/spectest/shared/common/forkchoice"
)

func TestMinimal_Bellatrix_Forkchoice(t *testing.T) {
	forkchoice.Run(t, "minimal", version.Bellatrix)
}
