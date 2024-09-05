package core

import (
	"github.com/Kevionte/prysm_beacon/v5/beacon-chain/blockchain"
	"github.com/Kevionte/prysm_beacon/v5/beacon-chain/cache"
	opfeed "github.com/Kevionte/prysm_beacon/v5/beacon-chain/core/feed/operation"
	"github.com/Kevionte/prysm_beacon/v5/beacon-chain/operations/synccommittee"
	"github.com/Kevionte/prysm_beacon/v5/beacon-chain/p2p"
	"github.com/Kevionte/prysm_beacon/v5/beacon-chain/state/stategen"
	"github.com/Kevionte/prysm_beacon/v5/beacon-chain/sync"
)

type Service struct {
	HeadFetcher           blockchain.HeadFetcher
	FinalizedFetcher      blockchain.FinalizationFetcher
	GenesisTimeFetcher    blockchain.TimeFetcher
	SyncChecker           sync.Checker
	Broadcaster           p2p.Broadcaster
	SyncCommitteePool     synccommittee.Pool
	OperationNotifier     opfeed.Notifier
	AttestationCache      *cache.AttestationCache
	StateGen              stategen.StateManager
	P2P                   p2p.Broadcaster
	OptimisticModeFetcher blockchain.OptimisticModeFetcher
}
