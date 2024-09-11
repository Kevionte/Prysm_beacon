package bazel_test

import (
	"testing"

	"github.com/Kevionte/prysm_beacon/v1build/bazel"
)

func TestBuildWithBazel(t *testing.T) {
	if !bazel.BuiltWithBazel() {
		t.Error("not built with Bazel")
	}
}
