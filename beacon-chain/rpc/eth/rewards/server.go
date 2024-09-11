package rewards

import (
	"github.com/Kevionte/prysm_beacon/v2/beacon-chain/blockchain"
	"github.com/Kevionte/prysm_beacon/v2/beacon-chain/rpc/lookup"
)

type Server struct {
	Blocker               lookup.Blocker
	OptimisticModeFetcher blockchain.OptimisticModeFetcher
	FinalizationFetcher   blockchain.FinalizationFetcher
	TimeFetcher           blockchain.TimeFetcher
	Stater                lookup.Stater
	HeadFetcher           blockchain.HeadFetcher
	BlockRewardFetcher    BlockRewardsFetcher
}
