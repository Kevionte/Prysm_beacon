package testing

import (
	"github.com/Kevionte/prysm_beacon/v2/time/slots"
)

var _ slots.Ticker = (*MockTicker)(nil)
