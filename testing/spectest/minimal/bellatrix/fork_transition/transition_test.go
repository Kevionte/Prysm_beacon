package fork_transition

import (
	"testing"

	"github.com/Kevionte/prysm_beacon/v1testing/spectest/shared/bellatrix/fork"
)

func TestMinimal_Bellatrix_Transition(t *testing.T) {
	fork.RunForkTransitionTest(t, "minimal")
}
