package operations

import (
	"context"
	"path"
	"testing"

	"github.com/Kevionte/prysm_beacon/v1beacon-chain/core/altair"
	"github.com/Kevionte/prysm_beacon/v1beacon-chain/state"
	"github.com/Kevionte/prysm_beacon/v1consensus-types/interfaces"
	ethpb "github.com/Kevionte/prysm_beacon/v1proto/prysm/v1alpha1"
	"github.com/Kevionte/prysm_beacon/v1testing/require"
	"github.com/Kevionte/prysm_beacon/v1testing/spectest/utils"
	"github.com/Kevionte/prysm_beacon/v1testing/util"
	"github.com/golang/snappy"
)

func RunSyncCommitteeTest(t *testing.T, config string) {
	require.NoError(t, utils.SetConfig(t, config))
	testFolders, testsFolderPath := utils.TestFolders(t, config, "deneb", "operations/sync_aggregate/pyspec_tests")
	for _, folder := range testFolders {
		t.Run(folder.Name(), func(t *testing.T) {
			folderPath := path.Join(testsFolderPath, folder.Name())
			syncCommitteeFile, err := util.BazelFileBytes(folderPath, "sync_aggregate.ssz_snappy")
			require.NoError(t, err)
			syncCommitteeSSZ, err := snappy.Decode(nil /* dst */, syncCommitteeFile)
			require.NoError(t, err, "Failed to decompress")
			sc := &ethpb.SyncAggregate{}
			require.NoError(t, sc.UnmarshalSSZ(syncCommitteeSSZ), "Failed to unmarshal")

			body := &ethpb.BeaconBlockBodyDeneb{SyncAggregate: sc}
			RunBlockOperationTest(t, folderPath, body, func(ctx context.Context, s state.BeaconState, b interfaces.SignedBeaconBlock) (state.BeaconState, error) {
				st, _, err := altair.ProcessSyncAggregate(context.Background(), s, body.SyncAggregate)
				if err != nil {
					return nil, err
				}
				return st, nil
			})
		})
	}
}
