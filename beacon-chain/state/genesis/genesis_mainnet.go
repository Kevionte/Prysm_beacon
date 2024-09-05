//go:build !noMainnetGenesis
// +build !noMainnetGenesis

package genesis

import (
	_ "embed"

	"github.com/Kevionte/prysm_beacon/v5/config/params"
)

var (
	//go:embed mainnet.ssz.snappy
	mainnetRawSSZCompressed []byte // 1.8Mb
)

func init() {
	embeddedStates[params.MainnetName] = &mainnetRawSSZCompressed
}
