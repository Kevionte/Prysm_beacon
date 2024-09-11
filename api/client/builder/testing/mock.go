package testing

import (
	"context"

	"github.com/Kevionte/prysm_beacon/v1api/client/builder"
	"github.com/Kevionte/prysm_beacon/v1consensus-types/interfaces"
	"github.com/Kevionte/prysm_beacon/v1consensus-types/primitives"
	"github.com/Kevionte/prysm_beacon/v1encoding/bytesutil"
	v1 "github.com/Kevionte/prysm_beacon/v1proto/engine/v1"
	ethpb "github.com/Kevionte/prysm_beacon/v1proto/prysm/v1alpha1"
)

// MockClient is a mock implementation of BuilderClient.
type MockClient struct {
	RegisteredVals map[[48]byte]bool
}

// NewClient creates a new, correctly initialized mock.
func NewClient() MockClient {
	return MockClient{RegisteredVals: map[[48]byte]bool{}}
}

// NodeURL --
func (MockClient) NodeURL() string {
	return ""
}

// GetHeader --
func (MockClient) GetHeader(_ context.Context, _ primitives.Slot, _ [32]byte, _ [48]byte) (builder.SignedBid, error) {
	return nil, nil
}

// RegisterValidator --
func (m MockClient) RegisterValidator(_ context.Context, svr []*ethpb.SignedValidatorRegistrationV1) error {
	for _, r := range svr {
		b := bytesutil.ToBytes48(r.Message.Pubkey)
		m.RegisteredVals[b] = true
	}
	return nil
}

// SubmitBlindedBlock --
func (MockClient) SubmitBlindedBlock(_ context.Context, _ interfaces.ReadOnlySignedBeaconBlock) (interfaces.ExecutionData, *v1.BlobsBundle, error) {
	return nil, nil, nil
}

// Status --
func (MockClient) Status(_ context.Context) error {
	return nil
}
