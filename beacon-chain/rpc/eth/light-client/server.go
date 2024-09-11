package lightclient

import (
	"github.com/Kevionte/prysm_beacon/v2/beacon-chain/blockchain"
	"github.com/Kevionte/prysm_beacon/v2/beacon-chain/rpc/lookup"
)

type Server struct {
	Blocker     lookup.Blocker
	Stater      lookup.Stater
	HeadFetcher blockchain.HeadFetcher
}
