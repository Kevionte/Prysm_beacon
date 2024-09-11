package testing

import (
	"context"

	"github.com/Kevionte/prysm_beacon/v1api/server/structs"
	"github.com/Kevionte/prysm_beacon/v1beacon-chain/state"
	"github.com/Kevionte/prysm_beacon/v1consensus-types/interfaces"
	"github.com/Kevionte/prysm_beacon/v1network/httputil"
)

type MockBlockRewardFetcher struct {
	Rewards *structs.BlockRewards
	Error   *httputil.DefaultJsonError
	State   state.BeaconState
}

func (m *MockBlockRewardFetcher) GetBlockRewardsData(_ context.Context, _ interfaces.ReadOnlyBeaconBlock) (*structs.BlockRewards, *httputil.DefaultJsonError) {
	if m.Error != nil {
		return nil, m.Error
	}
	return m.Rewards, nil
}

func (m *MockBlockRewardFetcher) GetStateForRewards(_ context.Context, _ interfaces.ReadOnlyBeaconBlock) (state.BeaconState, *httputil.DefaultJsonError) {
	if m.Error != nil {
		return nil, m.Error
	}
	return m.State, nil
}
