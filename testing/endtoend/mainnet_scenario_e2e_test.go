package endtoend

import (
	"testing"

	"github.com/Kevionte/prysm_beacon/v1config/params"
	"github.com/Kevionte/prysm_beacon/v1runtime/version"
	"github.com/Kevionte/prysm_beacon/v1testing/endtoend/types"
)

func TestEndToEnd_MultiScenarioRun_Multiclient(t *testing.T) {
	runner := e2eMainnet(t, false, true, types.InitForkCfg(version.Phase0, version.Deneb, params.E2EMainnetTestConfig()), types.WithEpochs(24))
	runner.config.Evaluators = scenarioEvalsMulti()
	runner.config.EvalInterceptor = runner.multiScenarioMulticlient
	runner.scenarioRunner()
}
