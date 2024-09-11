package checkpoint

import (
	"context"

	"github.com/Kevionte/prysm_beacon/v1api/client/beacon"
	"github.com/Kevionte/prysm_beacon/v1beacon-chain/db"
	"github.com/Kevionte/prysm_beacon/v1config/params"
	"github.com/pkg/errors"
)

// APIInitializer manages initializing the beacon node using checkpoint sync, retrieving the checkpoint state and root
// from the remote beacon node api.
type APIInitializer struct {
	c *beacon.Client
}

// NewAPIInitializer creates an APIInitializer, handling the set up of a beacon node api client
// using the provided host string.
func NewAPIInitializer(beaconNodeHost string) (*APIInitializer, error) {
	c, err := beacon.NewClient(beaconNodeHost)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to parse beacon node url or hostname - %s", beaconNodeHost)
	}
	return &APIInitializer{c: c}, nil
}

// Initialize downloads origin state and block for checkpoint sync and initializes database records to
// prepare the node to begin syncing from that point.
func (dl *APIInitializer) Initialize(ctx context.Context, d db.Database) error {
	origin, err := d.OriginCheckpointBlockRoot(ctx)
	if err == nil && origin != params.BeaconConfig().ZeroHash {
		log.Warnf("Origin checkpoint root %#x found in db, ignoring checkpoint sync flags", origin)
		return nil
	} else {
		if !errors.Is(err, db.ErrNotFound) {
			return errors.Wrap(err, "error while checking database for origin root")
		}
	}
	od, err := beacon.DownloadFinalizedData(ctx, dl.c)
	if err != nil {
		return errors.Wrap(err, "Error retrieving checkpoint origin state and block")
	}
	return d.SaveOrigin(ctx, od.StateBytes(), od.BlockBytes())
}
