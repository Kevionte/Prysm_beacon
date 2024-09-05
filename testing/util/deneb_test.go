package util

import (
	"testing"

	fieldparams "github.com/Kevionte/prysm_beacon/v5/config/fieldparams"
	"github.com/Kevionte/prysm_beacon/v5/consensus-types/blocks"
	"github.com/Kevionte/prysm_beacon/v5/testing/require"
)

func TestInclusionProofs(t *testing.T) {
	_, blobs := GenerateTestDenebBlockWithSidecar(t, [32]byte{}, 0, fieldparams.MaxBlobsPerBlock)
	for i := range blobs {
		require.NoError(t, blocks.VerifyKZGInclusionProof(blobs[i]))
	}
}
