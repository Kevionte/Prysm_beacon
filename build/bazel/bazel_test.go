package bazel_test

import (
	"testing"

	"github.com/Kevionte/prysm_beacon/v2/build/bazel"
)

func TestBuildWithBazel(t *testing.T) {
	if !bazel.BuiltWithBazel() {
		t.Error("not built with Bazel")
	}
}
