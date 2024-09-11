package epoch_processing

import (
	"path"
	"testing"

	"github.com/Kevionte/prysm_beacon/v1beacon-chain/core/epoch"
	"github.com/Kevionte/prysm_beacon/v1beacon-chain/state"
	"github.com/Kevionte/prysm_beacon/v1testing/require"
	"github.com/Kevionte/prysm_beacon/v1testing/spectest/utils"
)

// RunHistoricalRootsUpdateTests executes "epoch_processing/historical_roots_update" tests.
func RunHistoricalRootsUpdateTests(t *testing.T, config string) {
	require.NoError(t, utils.SetConfig(t, config))

	testFolders, testsFolderPath := utils.TestFolders(t, config, "bellatrix", "epoch_processing/historical_roots_update/pyspec_tests")
	if len(testFolders) == 0 {
		t.Fatalf("No test folders found for %s/%s/%s", config, "bellatrix", "epoch_processing/historical_roots_update/pyspec_tests")
	}
	for _, folder := range testFolders {
		t.Run(folder.Name(), func(t *testing.T) {
			folderPath := path.Join(testsFolderPath, folder.Name())
			RunEpochOperationTest(t, folderPath, processHistoricalRootsUpdateWrapper)
		})
	}
}

func processHistoricalRootsUpdateWrapper(t *testing.T, st state.BeaconState) (state.BeaconState, error) {
	st, err := epoch.ProcessHistoricalDataUpdate(st)
	require.NoError(t, err, "Could not process final updates")
	return st, nil
}
