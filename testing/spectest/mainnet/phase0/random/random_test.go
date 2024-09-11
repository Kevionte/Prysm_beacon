package random

import (
	"testing"

	"github.com/Kevionte/prysm_beacon/v1testing/spectest/shared/phase0/sanity"
)

func TestMainnet_Phase0_Random(t *testing.T) {
	sanity.RunBlockProcessingTest(t, "mainnet", "random/random/pyspec_tests")
}
