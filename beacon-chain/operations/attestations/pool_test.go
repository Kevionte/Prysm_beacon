package attestations

import (
	"github.com/Kevionte/prysm_beacon/v2/beacon-chain/operations/attestations/kv"
)

var _ Pool = (*kv.AttCaches)(nil)
