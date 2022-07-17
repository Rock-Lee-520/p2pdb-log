package event

import (
	"github.com/Rock-liyi/p2pdb/application/event"
)

func PublishSyncEvent(eventType string, data []byte) {
	event.PublishSyncEvent(eventType, data)
}
