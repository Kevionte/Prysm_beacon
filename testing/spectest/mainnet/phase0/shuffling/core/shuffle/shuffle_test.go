package shuffle

import (
	"testing"

	"github.com/Kevionte/prysm_beacon/v1testing/spectest/shared/phase0/shuffling/core/shuffle"
)

func TestMainnet_Phase0_Shuffling_Core_Shuffle(t *testing.T) {
	shuffle.RunShuffleTests(t, "mainnet")
}
