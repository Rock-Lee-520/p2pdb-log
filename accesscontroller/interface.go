package accesscontroller // import "github.com/Rock-liyi/p2pdb-log/accesscontroller"

import (
	"github.com/Rock-liyi/p2pdb-log/identityprovider"
)

type LogEntry interface {
	GetPayload() []byte
	GetIdentity() *identityprovider.Identity
}

type CanAppendAdditionalContext interface {
	GetLogEntries() []LogEntry
}

type Interface interface {
	CanAppend(LogEntry, identityprovider.Interface, CanAppendAdditionalContext) error
}
