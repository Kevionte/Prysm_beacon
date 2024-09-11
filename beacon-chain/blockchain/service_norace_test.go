package blockchain

import (
	"context"
	"io"
	"testing"

	testDB "github.com/Kevionte/prysm_beacon/v1beacon-chain/db/testing"
	"github.com/Kevionte/prysm_beacon/v1consensus-types/blocks"
	"github.com/Kevionte/prysm_beacon/v1testing/require"
	"github.com/Kevionte/prysm_beacon/v1testing/util"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetOutput(io.Discard)
}

func TestChainService_SaveHead_DataRace(t *testing.T) {
	beaconDB := testDB.SetupDB(t)
	s := &Service{
		cfg: &config{BeaconDB: beaconDB},
	}
	b, err := blocks.NewSignedBeaconBlock(util.NewBeaconBlock())
	st, _ := util.DeterministicGenesisState(t, 1)
	require.NoError(t, err)
	go func() {
		require.NoError(t, s.saveHead(context.Background(), [32]byte{}, b, st))
	}()
	require.NoError(t, s.saveHead(context.Background(), [32]byte{}, b, st))
}
