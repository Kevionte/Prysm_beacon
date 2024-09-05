package attestations

import (
	"github.com/Kevionte/prysm_beacon/v5/beacon-chain/operations/attestations/kv"
)

var _ Pool = (*kv.AttCaches)(nil)
