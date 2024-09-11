package testing

import (
	"github.com/Kevionte/prysm_beacon/v1/time/slots"
)

var _ slots.Ticker = (*MockTicker)(nil)
