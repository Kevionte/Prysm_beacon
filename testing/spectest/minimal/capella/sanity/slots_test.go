package sanity

import (
	"testing"

	"github.com/Kevionte/prysm_beacon/v1testing/spectest/shared/capella/sanity"
)

func TestMinimal_Capella_Sanity_Slots(t *testing.T) {
	sanity.RunSlotProcessingTests(t, "minimal")
}
