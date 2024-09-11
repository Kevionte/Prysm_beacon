package block

import "github.com/Kevionte/prysm_beacon/v1async/event"

// Notifier interface defines the methods of the service that provides block updates to consumers.
type Notifier interface {
	BlockFeed() *event.Feed
}
