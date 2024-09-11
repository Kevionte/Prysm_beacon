package epoch_processing

import (
	"path"
	"testing"

	"github.com/Kevionte/prysm_beacon/v1/beacon-chain/core/epoch"
	"github.com/Kevionte/prysm_beacon/v1/beacon-chain/state"
	"github.com/Kevionte/prysm_beacon/v1/testing/require"
	"github.com/Kevionte/prysm_beacon/v1/testing/spectest/utils"
)

// RunEffectiveBalanceUpdatesTests executes "epoch_processing/effective_balance_updates" tests.
func RunEffectiveBalanceUpdatesTests(t *testing.T, config string) {
	require.NoError(t, utils.SetConfig(t, config))

	testFolders, testsFolderPath := utils.TestFolders(t, config, "deneb", "epoch_processing/effective_balance_updates/pyspec_tests")
	for _, folder := range testFolders {
		t.Run(folder.Name(), func(t *testing.T) {
			folderPath := path.Join(testsFolderPath, folder.Name())
			RunEpochOperationTest(t, folderPath, processEffectiveBalanceUpdatesWrapper)
		})
	}
}

func processEffectiveBalanceUpdatesWrapper(t *testing.T, st state.BeaconState) (state.BeaconState, error) {
	st, err := epoch.ProcessEffectiveBalanceUpdates(st)
	require.NoError(t, err, "Could not process final updates")
	return st, nil
}
