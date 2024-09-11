package epoch_processing

import (
	"testing"

	"github.com/Kevionte/prysm_beacon/v1testing/spectest/shared/phase0/epoch_processing"
)

func TestMinimal_Phase0_EpochProcessing_Eth1DataReset(t *testing.T) {
	epoch_processing.RunEth1DataResetTests(t, "minimal")
}
