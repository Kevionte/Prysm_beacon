package iface

import (
	"context"

	"github.com/Kevionte/prysm_beacon/v1/api/client/beacon"
	ethpb "github.com/Kevionte/prysm_beacon/v1/proto/prysm/v1alpha1"
	"github.com/golang/protobuf/ptypes/empty"
)

type NodeClient interface {
	GetSyncStatus(ctx context.Context, in *empty.Empty) (*ethpb.SyncStatus, error)
	GetGenesis(ctx context.Context, in *empty.Empty) (*ethpb.Genesis, error)
	GetVersion(ctx context.Context, in *empty.Empty) (*ethpb.Version, error)
	ListPeers(ctx context.Context, in *empty.Empty) (*ethpb.Peers, error)
	HealthTracker() *beacon.NodeHealthTracker
}
