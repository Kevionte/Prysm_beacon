package epoch_processing

import (
	"path"
	"testing"

	"github.com/Kevionte/prysm_beacon/v5/beacon-chain/core/altair"
	"github.com/Kevionte/prysm_beacon/v5/beacon-chain/state"
	"github.com/Kevionte/prysm_beacon/v5/testing/require"
	"github.com/Kevionte/prysm_beacon/v5/testing/spectest/utils"
)

// RunParticipationFlagUpdatesTests executes "epoch_processing/participation_flag_updates" tests.
func RunParticipationFlagUpdatesTests(t *testing.T, config string) {
	require.NoError(t, utils.SetConfig(t, config))

	testFolders, testsFolderPath := utils.TestFolders(t, config, "deneb", "epoch_processing/participation_flag_updates/pyspec_tests")
	for _, folder := range testFolders {
		t.Run(folder.Name(), func(t *testing.T) {
			folderPath := path.Join(testsFolderPath, folder.Name())
			RunEpochOperationTest(t, folderPath, processParticipationFlagUpdatesWrapper)
		})
	}
}

func processParticipationFlagUpdatesWrapper(t *testing.T, st state.BeaconState) (state.BeaconState, error) {
	st, err := altair.ProcessParticipationFlagUpdates(st)
	require.NoError(t, err, "Could not process participation flag update")
	return st, nil
}
