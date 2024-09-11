package db

import "github.com/Kevionte/prysm_beacon/v1beacon-chain/db/kv"

var _ Database = (*kv.Store)(nil)
