package operations

import (
	"testing"

	"github.com/Kevionte/prysm_beacon/v2/testing/spectest/shared/deneb/operations"
)

func TestMainnet_Deneb_Operations_Attestation(t *testing.T) {
	operations.RunAttestationTest(t, "mainnet")
}
