package fork_transition

import (
	"testing"

	"github.com/Kevionte/prysm_beacon/v1testing/spectest/shared/capella/fork"
)

func TestMinimal_Capella_Transition(t *testing.T) {
	fork.RunForkTransitionTest(t, "minimal")
}
