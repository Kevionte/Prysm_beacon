package beacon

import (
	"github.com/Kevionte/prysm_beacon/v1beacon-chain/blockchain"
	beacondb "github.com/Kevionte/prysm_beacon/v1beacon-chain/db"
	"github.com/Kevionte/prysm_beacon/v1beacon-chain/rpc/lookup"
	"github.com/Kevionte/prysm_beacon/v1beacon-chain/state/stategen"
	"github.com/Kevionte/prysm_beacon/v1beacon-chain/sync"
)

type Server struct {
	SyncChecker           sync.Checker
	HeadFetcher           blockchain.HeadFetcher
	TimeFetcher           blockchain.TimeFetcher
	OptimisticModeFetcher blockchain.OptimisticModeFetcher
	CanonicalHistory      *stategen.CanonicalHistory
	BeaconDB              beacondb.ReadOnlyDatabase
	Stater                lookup.Stater
	ChainInfoFetcher      blockchain.ChainInfoFetcher
	FinalizationFetcher   blockchain.FinalizationFetcher
}
