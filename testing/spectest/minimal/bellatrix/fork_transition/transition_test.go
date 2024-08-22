package fork_transition

import (
	"testing"

	"github.com/prysmaticlabs/prysm/v5/testing/spectest/shared/bellatrix/fork"
)

func TestMinimal_Bellatrix_Transition(t *testing.T) {
	fork.RunForkTransitionTest(t, "minimal")
}
