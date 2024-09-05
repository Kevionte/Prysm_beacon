package epoch_processing

import (
	"testing"

	"github.com/Kevionte/prysm_beacon/v5/testing/spectest/shared/deneb/epoch_processing"
)

func TestMainnet_Deneb_EpochProcessing_HistoricalSummariesUpdate(t *testing.T) {
	epoch_processing.RunHistoricalSummariesUpdateTests(t, "mainnet")
}
