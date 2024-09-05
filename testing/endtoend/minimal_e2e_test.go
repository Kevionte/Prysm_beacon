package endtoend

import (
	"testing"

	"github.com/Kevionte/prysm_beacon/v5/config/params"
	"github.com/Kevionte/prysm_beacon/v5/runtime/version"
	"github.com/Kevionte/prysm_beacon/v5/testing/endtoend/types"
)

func TestEndToEnd_MinimalConfig(t *testing.T) {
	r := e2eMinimal(t, types.InitForkCfg(version.Phase0, version.Deneb, params.E2ETestConfig()), types.WithCheckpointSync())
	r.run()
}
