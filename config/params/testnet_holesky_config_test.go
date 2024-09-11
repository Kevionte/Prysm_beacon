package params_test

import (
	"path"
	"testing"

	"github.com/Kevionte/prysm_beacon/v1config/params"
	"github.com/Kevionte/prysm_beacon/v1testing/require"
	"github.com/bazelbuild/rules_go/go/tools/bazel"
)

func TestHoleskyConfigMatchesUpstreamYaml(t *testing.T) {
	presetFPs := presetsFilePath(t, "mainnet")
	mn, err := params.ByName(params.MainnetName)
	require.NoError(t, err)
	cfg := mn.Copy()
	for _, fp := range presetFPs {
		cfg, err = params.UnmarshalConfigFile(fp, cfg)
		require.NoError(t, err)
	}
	fPath, err := bazel.Runfile("external/holesky_testnet")
	require.NoError(t, err)
	configFP := path.Join(fPath, "custom_config_data", "config.yaml")
	pcfg, err := params.UnmarshalConfigFile(configFP, nil)
	require.NoError(t, err)
	fields := fieldsFromYamls(t, append(presetFPs, configFP))
	assertYamlFieldsMatch(t, "holesky", fields, pcfg, params.HoleskyConfig())
}
