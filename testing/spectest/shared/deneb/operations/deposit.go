package operations

import (
	"context"
	"path"
	"testing"

	"github.com/Kevionte/prysm_beacon/v1/beacon-chain/core/altair"
	"github.com/Kevionte/prysm_beacon/v1/beacon-chain/state"
	"github.com/Kevionte/prysm_beacon/v1/consensus-types/interfaces"
	ethpb "github.com/Kevionte/prysm_beacon/v1/proto/prysm/v1alpha1"
	"github.com/Kevionte/prysm_beacon/v1/testing/require"
	"github.com/Kevionte/prysm_beacon/v1/testing/spectest/utils"
	"github.com/Kevionte/prysm_beacon/v1/testing/util"
	"github.com/golang/snappy"
)

func RunDepositTest(t *testing.T, config string) {
	require.NoError(t, utils.SetConfig(t, config))
	testFolders, testsFolderPath := utils.TestFolders(t, config, "deneb", "operations/deposit/pyspec_tests")
	for _, folder := range testFolders {
		t.Run(folder.Name(), func(t *testing.T) {
			folderPath := path.Join(testsFolderPath, folder.Name())
			depositFile, err := util.BazelFileBytes(folderPath, "deposit.ssz_snappy")
			require.NoError(t, err)
			depositSSZ, err := snappy.Decode(nil /* dst */, depositFile)
			require.NoError(t, err, "Failed to decompress")
			deposit := &ethpb.Deposit{}
			require.NoError(t, deposit.UnmarshalSSZ(depositSSZ), "Failed to unmarshal")

			body := &ethpb.BeaconBlockBodyDeneb{Deposits: []*ethpb.Deposit{deposit}}
			processDepositsFunc := func(ctx context.Context, s state.BeaconState, b interfaces.SignedBeaconBlock) (state.BeaconState, error) {
				return altair.ProcessDeposits(ctx, s, b.Block().Body().Deposits())
			}
			RunBlockOperationTest(t, folderPath, body, processDepositsFunc)
		})
	}
}
