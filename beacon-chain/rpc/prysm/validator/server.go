package validator

import (
	"github.com/Kevionte/prysm_beacon/v2/beacon-chain/rpc/core"
)

type Server struct {
	CoreService *core.Service
}
