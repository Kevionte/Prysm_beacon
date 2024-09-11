package epoch_processing

import (
	"context"
	"path"
	"testing"

	"github.com/Kevionte/prysm_beacon/v1/beacon-chain/core/altair"
	"github.com/Kevionte/prysm_beacon/v1/beacon-chain/core/epoch/precompute"
	"github.com/Kevionte/prysm_beacon/v1/beacon-chain/state"
	"github.com/Kevionte/prysm_beacon/v1/testing/require"
	"github.com/Kevionte/prysm_beacon/v1/testing/spectest/utils"
)

// RunJustificationAndFinalizationTests executes "epoch_processing/justification_and_finalization" tests.
func RunJustificationAndFinalizationTests(t *testing.T, config string) {
	require.NoError(t, utils.SetConfig(t, config))

	testPath := "epoch_processing/justification_and_finalization/pyspec_tests"
	testFolders, testsFolderPath := utils.TestFolders(t, config, "deneb", testPath)
	for _, folder := range testFolders {
		t.Run(folder.Name(), func(t *testing.T) {
			folderPath := path.Join(testsFolderPath, folder.Name())
			RunEpochOperationTest(t, folderPath, processJustificationAndFinalizationPrecomputeWrapper)
		})
	}
}

func processJustificationAndFinalizationPrecomputeWrapper(t *testing.T, st state.BeaconState) (state.BeaconState, error) {
	ctx := context.Background()
	vp, bp, err := altair.InitializePrecomputeValidators(ctx, st)
	require.NoError(t, err)
	_, bp, err = altair.ProcessEpochParticipation(ctx, st, bp, vp)
	require.NoError(t, err)
	activeBal, targetPrevious, targetCurrent, err := st.UnrealizedCheckpointBalances()
	require.NoError(t, err)
	require.Equal(t, bp.ActiveCurrentEpoch, activeBal)
	require.Equal(t, bp.CurrentEpochTargetAttested, targetCurrent)
	require.Equal(t, bp.PrevEpochTargetAttested, targetPrevious)
	st, err = precompute.ProcessJustificationAndFinalizationPreCompute(st, bp)
	require.NoError(t, err, "Could not process justification")

	return st, nil
}
