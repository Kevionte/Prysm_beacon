package validator_client_factory

import (
	"github.com/Kevionte/prysm_beacon/v2/config/features"
	beaconApi "github.com/Kevionte/prysm_beacon/v2/validator/client/beacon-api"
	grpcApi "github.com/Kevionte/prysm_beacon/v2/validator/client/grpc-api"
	"github.com/Kevionte/prysm_beacon/v2/validator/client/iface"
	validatorHelpers "github.com/Kevionte/prysm_beacon/v2/validator/helpers"
)

func NewValidatorClient(
	validatorConn validatorHelpers.NodeConnection,
	jsonRestHandler beaconApi.JsonRestHandler,
	opt ...beaconApi.ValidatorClientOpt,
) iface.ValidatorClient {
	if features.Get().EnableBeaconRESTApi {
		return beaconApi.NewBeaconApiValidatorClient(jsonRestHandler, opt...)
	} else {
		return grpcApi.NewGrpcValidatorClient(validatorConn.GetGrpcClientConn())
	}
}
