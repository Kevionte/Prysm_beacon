// Package node defines a gRPC node service implementation, providing
// useful endpoints for checking a node's sync status, peer info,
// genesis data, and version information.
package node

import (
	"github.com/Kevionte/prysm_beacon/v5/beacon-chain/blockchain"
	"github.com/Kevionte/prysm_beacon/v5/beacon-chain/db"
	"github.com/Kevionte/prysm_beacon/v5/beacon-chain/execution"
	"github.com/Kevionte/prysm_beacon/v5/beacon-chain/p2p"
	"github.com/Kevionte/prysm_beacon/v5/beacon-chain/sync"
	"google.golang.org/grpc"
)

// Server defines a server implementation of the gRPC Node service,
// providing RPC endpoints for verifying a beacon node's sync status, genesis and
// version information.
type Server struct {
	SyncChecker               sync.Checker
	OptimisticModeFetcher     blockchain.OptimisticModeFetcher
	Server                    *grpc.Server
	BeaconDB                  db.ReadOnlyDatabase
	PeersFetcher              p2p.PeersProvider
	PeerManager               p2p.PeerManager
	MetadataProvider          p2p.MetadataProvider
	GenesisTimeFetcher        blockchain.TimeFetcher
	HeadFetcher               blockchain.HeadFetcher
	ExecutionChainInfoFetcher execution.ChainInfoFetcher
}
