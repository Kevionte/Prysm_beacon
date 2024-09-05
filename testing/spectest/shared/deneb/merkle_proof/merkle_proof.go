package merkle_proof

import (
	"testing"

	common "github.com/Kevionte/prysm_beacon/v5/testing/spectest/shared/common/merkle_proof"
	"github.com/Kevionte/prysm_beacon/v5/testing/spectest/shared/deneb/ssz_static"
)

func RunMerkleProofTests(t *testing.T, config string) {
	common.RunMerkleProofTests(t, config, "deneb", ssz_static.UnmarshalledSSZ)
}
