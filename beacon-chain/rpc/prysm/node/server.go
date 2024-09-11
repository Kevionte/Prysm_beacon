package node

import (
	"github.com/Kevionte/prysm_beacon/v1beacon-chain/blockchain"
	"github.com/Kevionte/prysm_beacon/v1beacon-chain/db"
	"github.com/Kevionte/prysm_beacon/v1beacon-chain/execution"
	"github.com/Kevionte/prysm_beacon/v1beacon-chain/p2p"
	"github.com/Kevionte/prysm_beacon/v1beacon-chain/sync"
)

type Server struct {
	SyncChecker               sync.Checker
	OptimisticModeFetcher     blockchain.OptimisticModeFetcher
	BeaconDB                  db.ReadOnlyDatabase
	PeersFetcher              p2p.PeersProvider
	PeerManager               p2p.PeerManager
	MetadataProvider          p2p.MetadataProvider
	GenesisTimeFetcher        blockchain.TimeFetcher
	HeadFetcher               blockchain.HeadFetcher
	ExecutionChainInfoFetcher execution.ChainInfoFetcher
}
