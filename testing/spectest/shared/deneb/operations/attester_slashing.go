package operations

import (
	"context"
	"path"
	"testing"

	"github.com/Kevionte/prysm_beacon/v5/beacon-chain/core/blocks"
	"github.com/Kevionte/prysm_beacon/v5/beacon-chain/core/validators"
	"github.com/Kevionte/prysm_beacon/v5/beacon-chain/state"
	"github.com/Kevionte/prysm_beacon/v5/consensus-types/interfaces"
	ethpb "github.com/Kevionte/prysm_beacon/v5/proto/prysm/v1alpha1"
	"github.com/Kevionte/prysm_beacon/v5/testing/require"
	"github.com/Kevionte/prysm_beacon/v5/testing/spectest/utils"
	"github.com/Kevionte/prysm_beacon/v5/testing/util"
	"github.com/golang/snappy"
)

func RunAttesterSlashingTest(t *testing.T, config string) {
	require.NoError(t, utils.SetConfig(t, config))
	testFolders, testsFolderPath := utils.TestFolders(t, config, "deneb", "operations/attester_slashing/pyspec_tests")
	for _, folder := range testFolders {
		t.Run(folder.Name(), func(t *testing.T) {
			folderPath := path.Join(testsFolderPath, folder.Name())
			attSlashingFile, err := util.BazelFileBytes(folderPath, "attester_slashing.ssz_snappy")
			require.NoError(t, err)
			attSlashingSSZ, err := snappy.Decode(nil /* dst */, attSlashingFile)
			require.NoError(t, err, "Failed to decompress")
			attSlashing := &ethpb.AttesterSlashing{}
			require.NoError(t, attSlashing.UnmarshalSSZ(attSlashingSSZ), "Failed to unmarshal")

			body := &ethpb.BeaconBlockBodyDeneb{AttesterSlashings: []*ethpb.AttesterSlashing{attSlashing}}
			RunBlockOperationTest(t, folderPath, body, func(ctx context.Context, s state.BeaconState, b interfaces.SignedBeaconBlock) (state.BeaconState, error) {
				return blocks.ProcessAttesterSlashings(ctx, s, b.Block().Body().AttesterSlashings(), validators.SlashValidator)
			})
		})
	}
}
