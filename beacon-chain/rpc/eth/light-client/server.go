package lightclient

import (
	"github.com/Kevionte/prysm_beacon/v1beacon-chain/blockchain"
	"github.com/Kevionte/prysm_beacon/v1beacon-chain/rpc/lookup"
)

type Server struct {
	Blocker     lookup.Blocker
	Stater      lookup.Stater
	HeadFetcher blockchain.HeadFetcher
}
