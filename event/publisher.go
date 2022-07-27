package event

import (
	"github.com/Rock-liyi/p2pdb/application/event"
)

func PublishSyncEvent(eventType string, data []byte) {
	var ms event.Message
	ms.Data = data
	ms.Type = eventType
	event.PublishSyncEvent(eventType, ms)
}
