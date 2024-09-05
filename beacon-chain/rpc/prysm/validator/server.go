package validator

import (
	"github.com/Kevionte/prysm_beacon/v5/beacon-chain/rpc/core"
)

type Server struct {
	CoreService *core.Service
}
