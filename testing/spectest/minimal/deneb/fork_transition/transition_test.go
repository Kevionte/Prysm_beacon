package fork_transition

import (
	"testing"

	"github.com/Kevionte/prysm_beacon/v5/testing/spectest/shared/deneb/fork"
)

func TestMinimal_Deneb_Transition(t *testing.T) {
	fork.RunForkTransitionTest(t, "minimal")
}
