package blob

import (
	"github.com/Kevionte/prysm_beacon/v5/beacon-chain/rpc/lookup"
)

type Server struct {
	Blocker lookup.Blocker
}
