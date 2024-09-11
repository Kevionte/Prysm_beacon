package validator

import (
	"github.com/Kevionte/prysm_beacon/v1beacon-chain/rpc/core"
)

type Server struct {
	CoreService *core.Service
}
