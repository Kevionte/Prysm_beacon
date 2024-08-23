package endtoend

// This file contains the dependencies required for github.com/Kevionte/Go-Sovereign/cmd/geth.
// Having these dependencies listed here helps go mod understand that these dependencies are
// necessary for end to end tests since we build go-ethereum binary for this test.
import (
	_ "github.com/Kevionte/Go-Sovereign/accounts"          // Required for go-ethereum e2e.
	_ "github.com/Kevionte/Go-Sovereign/accounts/keystore" // Required for go-ethereum e2e.
	_ "github.com/Kevionte/Go-Sovereign/cmd/utils"         // Required for go-ethereum e2e.
	_ "github.com/Kevionte/Go-Sovereign/common"            // Required for go-ethereum e2e.
	_ "github.com/Kevionte/Go-Sovereign/console"           // Required for go-ethereum e2e.
	_ "github.com/Kevionte/Go-Sovereign/eth"               // Required for go-ethereum e2e.
	_ "github.com/Kevionte/Go-Sovereign/eth/downloader"    // Required for go-ethereum e2e.
	_ "github.com/Kevionte/Go-Sovereign/ethclient"         // Required for go-ethereum e2e.
	_ "github.com/Kevionte/Go-Sovereign/les"               // Required for go-ethereum e2e.
	_ "github.com/Kevionte/Go-Sovereign/log"               // Required for go-ethereum e2e.
	_ "github.com/Kevionte/Go-Sovereign/metrics"           // Required for go-ethereum e2e.
	_ "github.com/Kevionte/Go-Sovereign/node"              // Required for go-ethereum e2e.
)
