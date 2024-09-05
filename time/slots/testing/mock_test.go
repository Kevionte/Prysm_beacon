package testing

import (
	"github.com/Kevionte/prysm_beacon/v5/time/slots"
)

var _ slots.Ticker = (*MockTicker)(nil)
