package beacon_api

import (
	"bytes"
	"context"
	"encoding/json"
	"testing"

	"github.com/Kevionte/prysm_beacon/v1api/server/structs"
	rpctesting "github.com/Kevionte/prysm_beacon/v1beacon-chain/rpc/eth/shared/testing"
	"github.com/Kevionte/prysm_beacon/v1testing/assert"
	"github.com/Kevionte/prysm_beacon/v1testing/require"
	"github.com/Kevionte/prysm_beacon/v1validator/client/beacon-api/mock"
	"go.uber.org/mock/gomock"
)

func TestProposeBeaconBlock_BlindedDeneb(t *testing.T) {
	t.Skip("TODO: Fix this in the beacon-API PR")
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	jsonRestHandler := mock.NewMockJsonRestHandler(ctrl)

	var block structs.SignedBlindedBeaconBlockDeneb
	err := json.Unmarshal([]byte(rpctesting.BlindedDenebBlock), &block)
	require.NoError(t, err)
	genericSignedBlock, err := block.ToGeneric()
	require.NoError(t, err)

	denebBytes, err := json.Marshal(block)
	require.NoError(t, err)
	// Make sure that what we send in the POST body is the marshalled version of the protobuf block
	headers := map[string]string{"Eth-Consensus-Version": "deneb"}
	jsonRestHandler.EXPECT().Post(
		context.Background(),
		"/eth/v1/beacon/blinded_blocks",
		headers,
		bytes.NewBuffer(denebBytes),
		nil,
	)

	validatorClient := &beaconApiValidatorClient{jsonRestHandler: jsonRestHandler}
	proposeResponse, err := validatorClient.proposeBeaconBlock(context.Background(), genericSignedBlock)
	assert.NoError(t, err)
	require.NotNil(t, proposeResponse)

	expectedBlockRoot, err := genericSignedBlock.GetBlindedDeneb().HashTreeRoot()
	require.NoError(t, err)

	// Make sure that the block root is set
	assert.DeepEqual(t, expectedBlockRoot[:], proposeResponse.BlockRoot)
}
