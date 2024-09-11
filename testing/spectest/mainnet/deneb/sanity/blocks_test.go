package sanity

import (
	"testing"

	"github.com/Kevionte/prysm_beacon/v1testing/spectest/shared/deneb/sanity"
)

func TestMainnet_Deneb_Sanity_Blocks(t *testing.T) {
	sanity.RunBlockProcessingTest(t, "mainnet", "sanity/blocks/pyspec_tests")
}
