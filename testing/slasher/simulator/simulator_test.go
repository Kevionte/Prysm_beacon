package simulator

import (
	"testing"

	mock "github.com/Kevionte/prysm_beacon/v1beacon-chain/blockchain/testing"
	dbtest "github.com/Kevionte/prysm_beacon/v1beacon-chain/db/testing"
	mockstategen "github.com/Kevionte/prysm_beacon/v1beacon-chain/state/stategen/mock"
	"github.com/Kevionte/prysm_beacon/v1consensus-types/primitives"
	"github.com/Kevionte/prysm_beacon/v1crypto/bls"
	ethpb "github.com/Kevionte/prysm_beacon/v1proto/prysm/v1alpha1"
	"github.com/Kevionte/prysm_beacon/v1testing/require"
	"github.com/Kevionte/prysm_beacon/v1testing/util"
)

func setupService(t *testing.T, params *Parameters) *Simulator {
	slasherDB := dbtest.SetupSlasherDB(t)
	beaconState, err := util.NewBeaconState()
	require.NoError(t, err)

	// We setup validators in the beacon state along with their
	// private keys used to generate valid signatures in generated objects.
	validators := make([]*ethpb.Validator, params.NumValidators)
	privKeys := make(map[primitives.ValidatorIndex]bls.SecretKey)
	for valIdx := range validators {
		privKey, err := bls.RandKey()
		require.NoError(t, err)
		privKeys[primitives.ValidatorIndex(valIdx)] = privKey
		validators[valIdx] = &ethpb.Validator{
			PublicKey:             privKey.PublicKey().Marshal(),
			WithdrawalCredentials: make([]byte, 32),
		}
	}
	err = beaconState.SetValidators(validators)
	require.NoError(t, err)

	gen := mockstategen.NewService()
	gen.AddStateForRoot(beaconState, [32]byte{})
	return &Simulator{
		srvConfig: &ServiceConfig{
			Params:                      params,
			Database:                    slasherDB,
			AttestationStateFetcher:     &mock.ChainService{State: beaconState},
			PrivateKeysByValidatorIndex: privKeys,
			StateGen:                    gen,
		},
	}
}
