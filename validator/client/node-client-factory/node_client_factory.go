package validator_client_factory

import (
	"github.com/Kevionte/prysm_beacon/v2/config/features"
	beaconApi "github.com/Kevionte/prysm_beacon/v2/validator/client/beacon-api"
	grpcApi "github.com/Kevionte/prysm_beacon/v2/validator/client/grpc-api"
	"github.com/Kevionte/prysm_beacon/v2/validator/client/iface"
	validatorHelpers "github.com/Kevionte/prysm_beacon/v2/validator/helpers"
)

func NewNodeClient(validatorConn validatorHelpers.NodeConnection, jsonRestHandler beaconApi.JsonRestHandler) iface.NodeClient {
	grpcClient := grpcApi.NewNodeClient(validatorConn.GetGrpcClientConn())
	if features.Get().EnableBeaconRESTApi {
		return beaconApi.NewNodeClientWithFallback(jsonRestHandler, grpcClient)
	} else {
		return grpcClient
	}
}
