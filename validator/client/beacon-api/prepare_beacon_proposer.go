package beacon_api

import (
	"bytes"
	"context"
	"encoding/json"
	"strconv"

	"github.com/Kevionte/go-sovereign/common/hexutil"
	"github.com/Kevionte/prysm_beacon/v1api/server/structs"
	ethpb "github.com/Kevionte/prysm_beacon/v1proto/prysm/v1alpha1"
	"github.com/pkg/errors"
)

func (c *beaconApiValidatorClient) prepareBeaconProposer(ctx context.Context, recipients []*ethpb.PrepareBeaconProposerRequest_FeeRecipientContainer) error {
	jsonRecipients := make([]*structs.FeeRecipient, len(recipients))
	for index, recipient := range recipients {
		jsonRecipients[index] = &structs.FeeRecipient{
			FeeRecipient:   hexutil.Encode(recipient.FeeRecipient),
			ValidatorIndex: strconv.FormatUint(uint64(recipient.ValidatorIndex), 10),
		}
	}

	marshalledJsonRecipients, err := json.Marshal(jsonRecipients)
	if err != nil {
		return errors.Wrap(err, "failed to marshal recipients")
	}

	return c.jsonRestHandler.Post(ctx, "/eth/v1/validator/prepare_beacon_proposer", nil, bytes.NewBuffer(marshalledJsonRecipients), nil)
}
