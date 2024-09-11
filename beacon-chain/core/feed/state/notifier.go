package state

import "github.com/Kevionte/prysm_beacon/v2/async/event"

// Notifier interface defines the methods of the service that provides state updates to consumers.
type Notifier interface {
	StateFeed() *event.Feed
}
