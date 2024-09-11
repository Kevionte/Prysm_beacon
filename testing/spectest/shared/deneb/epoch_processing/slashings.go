package epoch_processing

import (
	"path"
	"testing"

	"github.com/Kevionte/prysm_beacon/v1beacon-chain/core/epoch"
	"github.com/Kevionte/prysm_beacon/v1beacon-chain/core/helpers"
	"github.com/Kevionte/prysm_beacon/v1beacon-chain/state"
	"github.com/Kevionte/prysm_beacon/v1config/params"
	"github.com/Kevionte/prysm_beacon/v1testing/require"
	"github.com/Kevionte/prysm_beacon/v1testing/spectest/utils"
)

// RunSlashingsTests executes "epoch_processing/slashings" tests.
func RunSlashingsTests(t *testing.T, config string) {
	require.NoError(t, utils.SetConfig(t, config))

	testFolders, testsFolderPath := utils.TestFolders(t, config, "deneb", "epoch_processing/slashings/pyspec_tests")
	for _, folder := range testFolders {
		t.Run(folder.Name(), func(t *testing.T) {
			folderPath := path.Join(testsFolderPath, folder.Name())
			helpers.ClearCache()
			RunEpochOperationTest(t, folderPath, processSlashingsWrapper)
		})
	}
}

func processSlashingsWrapper(t *testing.T, st state.BeaconState) (state.BeaconState, error) {
	st, err := epoch.ProcessSlashings(st, params.BeaconConfig().ProportionalSlashingMultiplierBellatrix)
	require.NoError(t, err, "Could not process slashings")
	return st, nil
}
