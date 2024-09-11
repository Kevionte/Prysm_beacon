// Package beacon defines a gRPC beacon service implementation, providing
// useful endpoints for checking fetching chain-specific data such as
// blocks, committees, validators, assignments, and more.
package beacon

import (
	"context"
	"time"

	"github.com/Kevionte/prysm_beacon/v1beacon-chain/blockchain"
	"github.com/Kevionte/prysm_beacon/v1beacon-chain/cache"
	blockfeed "github.com/Kevionte/prysm_beacon/v1beacon-chain/core/feed/block"
	"github.com/Kevionte/prysm_beacon/v1beacon-chain/core/feed/operation"
	statefeed "github.com/Kevionte/prysm_beacon/v1beacon-chain/core/feed/state"
	"github.com/Kevionte/prysm_beacon/v1beacon-chain/db"
	"github.com/Kevionte/prysm_beacon/v1beacon-chain/execution"
	"github.com/Kevionte/prysm_beacon/v1beacon-chain/operations/attestations"
	"github.com/Kevionte/prysm_beacon/v1beacon-chain/operations/slashings"
	"github.com/Kevionte/prysm_beacon/v1beacon-chain/p2p"
	"github.com/Kevionte/prysm_beacon/v1beacon-chain/rpc/core"
	"github.com/Kevionte/prysm_beacon/v1beacon-chain/state/stategen"
	"github.com/Kevionte/prysm_beacon/v1beacon-chain/sync"
	ethpb "github.com/Kevionte/prysm_beacon/v1proto/prysm/v1alpha1"
)

// Server defines a server implementation of the gRPC Beacon Chain service,
// providing RPC endpoints to access data relevant to the Ethereum beacon chain.
type Server struct {
	BeaconDB                    db.ReadOnlyDatabase
	Ctx                         context.Context
	ChainStartFetcher           execution.ChainStartFetcher
	HeadFetcher                 blockchain.HeadFetcher
	CanonicalFetcher            blockchain.CanonicalFetcher
	FinalizationFetcher         blockchain.FinalizationFetcher
	DepositFetcher              cache.DepositFetcher
	BlockFetcher                execution.POWBlockFetcher
	GenesisTimeFetcher          blockchain.TimeFetcher
	StateNotifier               statefeed.Notifier
	BlockNotifier               blockfeed.Notifier
	AttestationNotifier         operation.Notifier
	Broadcaster                 p2p.Broadcaster
	AttestationsPool            attestations.Pool
	SlashingsPool               slashings.PoolManager
	ChainStartChan              chan time.Time
	ReceivedAttestationsBuffer  chan *ethpb.Attestation
	CollectedAttestationsBuffer chan []*ethpb.Attestation
	StateGen                    stategen.StateManager
	SyncChecker                 sync.Checker
	ReplayerBuilder             stategen.ReplayerBuilder
	OptimisticModeFetcher       blockchain.OptimisticModeFetcher
	CoreService                 *core.Service
}
