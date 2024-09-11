package random

import (
	"testing"

	"github.com/Kevionte/prysm_beacon/v1testing/spectest/shared/capella/sanity"
)

func TestMainnet_Capella_Random(t *testing.T) {
	sanity.RunBlockProcessingTest(t, "mainnet", "random/random/pyspec_tests")
}
