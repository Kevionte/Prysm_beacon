// Package beacon defines a gRPC beacon service implementation,
// following the official API standards https://ethereum.github.io/beacon-apis/#/.
// This package includes the beacon and config endpoints.
package beacon

import (
	"github.com/Kevionte/prysm_beacon/v2/beacon-chain/blockchain"
	blockfeed "github.com/Kevionte/prysm_beacon/v2/beacon-chain/core/feed/block"
	"github.com/Kevionte/prysm_beacon/v2/beacon-chain/core/feed/operation"
	"github.com/Kevionte/prysm_beacon/v2/beacon-chain/db"
	"github.com/Kevionte/prysm_beacon/v2/beacon-chain/execution"
	"github.com/Kevionte/prysm_beacon/v2/beacon-chain/operations/attestations"
	"github.com/Kevionte/prysm_beacon/v2/beacon-chain/operations/blstoexec"
	"github.com/Kevionte/prysm_beacon/v2/beacon-chain/operations/slashings"
	"github.com/Kevionte/prysm_beacon/v2/beacon-chain/operations/voluntaryexits"
	"github.com/Kevionte/prysm_beacon/v2/beacon-chain/p2p"
	"github.com/Kevionte/prysm_beacon/v2/beacon-chain/rpc/core"
	"github.com/Kevionte/prysm_beacon/v2/beacon-chain/rpc/lookup"
	"github.com/Kevionte/prysm_beacon/v2/beacon-chain/state/stategen"
	"github.com/Kevionte/prysm_beacon/v2/beacon-chain/sync"
	eth "github.com/Kevionte/prysm_beacon/v2/proto/prysm/v1alpha1"
)

// Server defines a server implementation of the gRPC Beacon Chain service,
// providing RPC endpoints to access data relevant to the Ethereum Beacon Chain.
type Server struct {
	BeaconDB                      db.ReadOnlyDatabase
	ChainInfoFetcher              blockchain.ChainInfoFetcher
	GenesisTimeFetcher            blockchain.TimeFetcher
	BlockReceiver                 blockchain.BlockReceiver
	BlockNotifier                 blockfeed.Notifier
	OperationNotifier             operation.Notifier
	Broadcaster                   p2p.Broadcaster
	AttestationsPool              attestations.Pool
	SlashingsPool                 slashings.PoolManager
	VoluntaryExitsPool            voluntaryexits.PoolManager
	StateGenService               stategen.StateManager
	Stater                        lookup.Stater
	Blocker                       lookup.Blocker
	HeadFetcher                   blockchain.HeadFetcher
	TimeFetcher                   blockchain.TimeFetcher
	OptimisticModeFetcher         blockchain.OptimisticModeFetcher
	V1Alpha1ValidatorServer       eth.BeaconNodeValidatorServer
	SyncChecker                   sync.Checker
	CanonicalHistory              *stategen.CanonicalHistory
	ExecutionPayloadReconstructor execution.PayloadReconstructor
	FinalizationFetcher           blockchain.FinalizationFetcher
	BLSChangesPool                blstoexec.PoolManager
	ForkchoiceFetcher             blockchain.ForkchoiceFetcher
	CoreService                   *core.Service
}
