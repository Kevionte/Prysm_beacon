package endtoend

// This file contains the dependencies required for github.com/Kevionte/go-sovereign/cmd/geth.
// Having these dependencies listed here helps go mod understand that these dependencies are
// necessary for end to end tests since we build go-ethereum binary for this test.
import (
	_ "github.com/Kevionte/go-sovereign/accounts"          // Required for go-ethereum e2e.
	_ "github.com/Kevionte/go-sovereign/accounts/keystore" // Required for go-ethereum e2e.
	_ "github.com/Kevionte/go-sovereign/cmd/utils"         // Required for go-ethereum e2e.
	_ "github.com/Kevionte/go-sovereign/common"            // Required for go-ethereum e2e.
	_ "github.com/Kevionte/go-sovereign/console"           // Required for go-ethereum e2e.
	_ "github.com/Kevionte/go-sovereign/eth"               // Required for go-ethereum e2e.
	_ "github.com/Kevionte/go-sovereign/eth/downloader"    // Required for go-ethereum e2e.
	_ "github.com/Kevionte/go-sovereign/ethclient"         // Required for go-ethereum e2e.
	_ "github.com/Kevionte/go-sovereign/les"               // Required for go-ethereum e2e.
	_ "github.com/Kevionte/go-sovereign/log"               // Required for go-ethereum e2e.
	_ "github.com/Kevionte/go-sovereign/metrics"           // Required for go-ethereum e2e.
	_ "github.com/Kevionte/go-sovereign/node"              // Required for go-ethereum e2e.
)
