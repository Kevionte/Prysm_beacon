package operations

import (
	"context"
	"path"
	"testing"

	"github.com/Kevionte/prysm_beacon/v2/beacon-chain/core/blocks"
	"github.com/Kevionte/prysm_beacon/v2/beacon-chain/core/validators"
	"github.com/Kevionte/prysm_beacon/v2/beacon-chain/state"
	"github.com/Kevionte/prysm_beacon/v2/consensus-types/interfaces"
	ethpb "github.com/Kevionte/prysm_beacon/v2/proto/prysm/v1alpha1"
	"github.com/Kevionte/prysm_beacon/v2/testing/require"
	"github.com/Kevionte/prysm_beacon/v2/testing/spectest/utils"
	"github.com/Kevionte/prysm_beacon/v2/testing/util"
	"github.com/golang/snappy"
)

func RunProposerSlashingTest(t *testing.T, config string) {
	require.NoError(t, utils.SetConfig(t, config))
	testFolders, testsFolderPath := utils.TestFolders(t, config, "altair", "operations/proposer_slashing/pyspec_tests")
	if len(testFolders) == 0 {
		t.Fatalf("No test folders found for %s/%s/%s", config, "altair", "operations/proposer_slashing/pyspec_tests")
	}
	for _, folder := range testFolders {
		t.Run(folder.Name(), func(t *testing.T) {
			folderPath := path.Join(testsFolderPath, folder.Name())
			proposerSlashingFile, err := util.BazelFileBytes(folderPath, "proposer_slashing.ssz_snappy")
			require.NoError(t, err)
			proposerSlashingSSZ, err := snappy.Decode(nil /* dst */, proposerSlashingFile)
			require.NoError(t, err, "Failed to decompress")
			proposerSlashing := &ethpb.ProposerSlashing{}
			require.NoError(t, proposerSlashing.UnmarshalSSZ(proposerSlashingSSZ), "Failed to unmarshal")

			body := &ethpb.BeaconBlockBodyAltair{ProposerSlashings: []*ethpb.ProposerSlashing{proposerSlashing}}
			RunBlockOperationTest(t, folderPath, body, func(ctx context.Context, s state.BeaconState, b interfaces.ReadOnlySignedBeaconBlock) (state.BeaconState, error) {
				return blocks.ProcessProposerSlashings(ctx, s, b.Block().Body().ProposerSlashings(), validators.SlashValidator)
			})
		})
	}
}
