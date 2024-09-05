package builder

import (
	"github.com/Kevionte/prysm_beacon/v5/beacon-chain/blockchain"
	"github.com/Kevionte/prysm_beacon/v5/beacon-chain/rpc/lookup"
)

type Server struct {
	FinalizationFetcher   blockchain.FinalizationFetcher
	OptimisticModeFetcher blockchain.OptimisticModeFetcher
	Stater                lookup.Stater
}
