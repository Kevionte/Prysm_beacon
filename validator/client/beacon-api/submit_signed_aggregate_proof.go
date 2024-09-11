package beacon_api

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/Kevionte/prysm_beacon/v2/api/server/structs"
	ethpb "github.com/Kevionte/prysm_beacon/v2/proto/prysm/v1alpha1"
	"github.com/pkg/errors"
)

func (c *beaconApiValidatorClient) submitSignedAggregateSelectionProof(ctx context.Context, in *ethpb.SignedAggregateSubmitRequest) (*ethpb.SignedAggregateSubmitResponse, error) {
	body, err := json.Marshal([]*structs.SignedAggregateAttestationAndProof{jsonifySignedAggregateAndProof(in.SignedAggregateAndProof)})
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal SignedAggregateAttestationAndProof")
	}

	if err = c.jsonRestHandler.Post(ctx, "/eth/v1/validator/aggregate_and_proofs", nil, bytes.NewBuffer(body), nil); err != nil {
		return nil, err
	}

	attestationDataRoot, err := in.SignedAggregateAndProof.Message.Aggregate.Data.HashTreeRoot()
	if err != nil {
		return nil, errors.Wrap(err, "failed to compute attestation data root")
	}

	return &ethpb.SignedAggregateSubmitResponse{AttestationDataRoot: attestationDataRoot[:]}, nil
}
