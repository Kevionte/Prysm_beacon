package validator

import (
	"github.com/Kevionte/prysm_beacon/v1beacon-chain/blockchain"
	"github.com/Kevionte/prysm_beacon/v1beacon-chain/builder"
	"github.com/Kevionte/prysm_beacon/v1beacon-chain/cache"
	"github.com/Kevionte/prysm_beacon/v1beacon-chain/core/feed/operation"
	"github.com/Kevionte/prysm_beacon/v1beacon-chain/db"
	"github.com/Kevionte/prysm_beacon/v1beacon-chain/operations/attestations"
	"github.com/Kevionte/prysm_beacon/v1beacon-chain/operations/synccommittee"
	"github.com/Kevionte/prysm_beacon/v1beacon-chain/p2p"
	"github.com/Kevionte/prysm_beacon/v1beacon-chain/rpc/core"
	"github.com/Kevionte/prysm_beacon/v1beacon-chain/rpc/eth/rewards"
	"github.com/Kevionte/prysm_beacon/v1beacon-chain/rpc/lookup"
	"github.com/Kevionte/prysm_beacon/v1beacon-chain/sync"
	eth "github.com/Kevionte/prysm_beacon/v1proto/prysm/v1alpha1"
)

// Server defines a server implementation of the gRPC Validator service,
// providing RPC endpoints intended for validator clients.
type Server struct {
	HeadFetcher            blockchain.HeadFetcher
	TimeFetcher            blockchain.TimeFetcher
	SyncChecker            sync.Checker
	AttestationsPool       attestations.Pool
	PeerManager            p2p.PeerManager
	Broadcaster            p2p.Broadcaster
	Stater                 lookup.Stater
	OptimisticModeFetcher  blockchain.OptimisticModeFetcher
	SyncCommitteePool      synccommittee.Pool
	V1Alpha1Server         eth.BeaconNodeValidatorServer
	ChainInfoFetcher       blockchain.ChainInfoFetcher
	BeaconDB               db.HeadAccessDatabase
	BlockBuilder           builder.BlockBuilder
	OperationNotifier      operation.Notifier
	CoreService            *core.Service
	BlockRewardFetcher     rewards.BlockRewardsFetcher
	TrackedValidatorsCache *cache.TrackedValidatorsCache
	PayloadIDCache         *cache.PayloadIDCache
}
