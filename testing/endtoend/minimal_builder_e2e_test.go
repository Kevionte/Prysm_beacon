package endtoend

import (
	"testing"

	"github.com/Kevionte/prysm_beacon/v2/config/params"
	"github.com/Kevionte/prysm_beacon/v2/runtime/version"
	"github.com/Kevionte/prysm_beacon/v2/testing/endtoend/types"
)

func TestEndToEnd_MinimalConfig_WithBuilder(t *testing.T) {
	r := e2eMinimal(t, types.InitForkCfg(version.Phase0, version.Deneb, params.E2ETestConfig()), types.WithCheckpointSync(), types.WithBuilder())
	r.run()
}

func TestEndToEnd_MinimalConfig_WithBuilder_ValidatorRESTApi(t *testing.T) {
	r := e2eMinimal(t, types.InitForkCfg(version.Phase0, version.Deneb, params.E2ETestConfig()), types.WithCheckpointSync(), types.WithBuilder(), types.WithValidatorRESTApi())
	r.run()
}
