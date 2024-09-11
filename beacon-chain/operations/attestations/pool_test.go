package attestations

import (
	"github.com/Kevionte/prysm_beacon/v1beacon-chain/operations/attestations/kv"
)

var _ Pool = (*kv.AttCaches)(nil)
