package sanity

import (
	"context"
	"strconv"
	"testing"

	"github.com/Kevionte/prysm_beacon/v5/beacon-chain/core/transition"
	state_native "github.com/Kevionte/prysm_beacon/v5/beacon-chain/state/state-native"
	ethpb "github.com/Kevionte/prysm_beacon/v5/proto/prysm/v1alpha1"
	"github.com/Kevionte/prysm_beacon/v5/testing/require"
	"github.com/Kevionte/prysm_beacon/v5/testing/spectest/utils"
	"github.com/Kevionte/prysm_beacon/v5/testing/util"
	"github.com/golang/snappy"
	"google.golang.org/protobuf/proto"
	"gopkg.in/d4l3k/messagediff.v1"
)

func init() {
	transition.SkipSlotCache.Disable()
}

// RunSlotProcessingTests executes "sanity/slots" tests.
func RunSlotProcessingTests(t *testing.T, config string) {
	require.NoError(t, utils.SetConfig(t, config))

	testFolders, testsFolderPath := utils.TestFolders(t, config, "bellatrix", "sanity/slots/pyspec_tests")
	if len(testFolders) == 0 {
		t.Fatalf("No test folders found for %s/%s/%s", config, "bellatrix", "sanity/slots/pyspec_tests")
	}
	for _, folder := range testFolders {
		t.Run(folder.Name(), func(t *testing.T) {
			preBeaconStateFile, err := util.BazelFileBytes(testsFolderPath, folder.Name(), "pre.ssz_snappy")
			require.NoError(t, err)
			preBeaconStateSSZ, err := snappy.Decode(nil /* dst */, preBeaconStateFile)
			require.NoError(t, err, "Failed to decompress")
			base := &ethpb.BeaconStateBellatrix{}
			require.NoError(t, base.UnmarshalSSZ(preBeaconStateSSZ), "Failed to unmarshal")
			beaconState, err := state_native.InitializeFromProtoBellatrix(base)
			require.NoError(t, err)

			file, err := util.BazelFileBytes(testsFolderPath, folder.Name(), "slots.yaml")
			require.NoError(t, err)
			fileStr := string(file)
			slotsCount, err := strconv.ParseUint(fileStr[:len(fileStr)-5], 10, 64)
			require.NoError(t, err)

			postBeaconStateFile, err := util.BazelFileBytes(testsFolderPath, folder.Name(), "post.ssz_snappy")
			require.NoError(t, err)
			postBeaconStateSSZ, err := snappy.Decode(nil /* dst */, postBeaconStateFile)
			require.NoError(t, err, "Failed to decompress")
			postBeaconState := &ethpb.BeaconStateBellatrix{}
			require.NoError(t, postBeaconState.UnmarshalSSZ(postBeaconStateSSZ), "Failed to unmarshal")
			postState, err := transition.ProcessSlots(context.Background(), beaconState, beaconState.Slot().Add(slotsCount))
			require.NoError(t, err)

			pbState, err := state_native.ProtobufBeaconStateBellatrix(postState.ToProto())
			require.NoError(t, err)
			if !proto.Equal(pbState, postBeaconState) {
				diff, _ := messagediff.PrettyDiff(beaconState, postBeaconState)
				t.Fatalf("Post state does not match expected. Diff between states %s", diff)
			}
		})
	}
}
