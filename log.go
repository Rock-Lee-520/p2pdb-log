package main

import (
	"p2pdb-log/clock"
	"p2pdb-log/store"
	"sync"

	"berty.tech/go-ipfs-log/accesscontroller"
	"berty.tech/go-ipfs-log/entry"
	"berty.tech/go-ipfs-log/entry/sorting"
	"berty.tech/go-ipfs-log/errmsg"
	"berty.tech/go-ipfs-log/identityprovider"
	"berty.tech/go-ipfs-log/iface"
	core_iface "github.com/ipfs/interface-go-ipfs-core"
)

type Snapshot = iface.Snapshot
type JSONLog = iface.JSONLog
type LogOptions = iface.LogOptions
type IteratorOptions = iface.IteratorOptions
type IO = iface.IO

type Entry = iface.IPFSLogEntry

// type Log = iface.IPFSLog
type AppendOptions = iface.AppendOptions
type SortFn = iface.EntrySortFn

type Log interface {
	NewLog(services string, identity string, options string) (*P2PDBLog, error)
	Append()
	Join()
	ToJson()
	Pull()
	Push()
}

type P2PDBLog struct {
	Storage          core_iface.CoreAPI
	ID               string
	AccessController accesscontroller.Interface
	SortFn           iface.EntrySortFn
	Identity         *identityprovider.Identity
	Entries          iface.IPFSLogOrderedEntries
	heads            iface.IPFSLogOrderedEntries
	Next             iface.IPFSLogOrderedEntries
	lamportClock     iface.IPFSLogLamportClock
	io               iface.IO
	concurrency      uint
	lock             sync.RWMutex
}

func init() {

	var db *store.CreateDBFactory
	orm := db.InitDB()
	// init node、link、object table schema
	orm.Migrator().AutoMigrate(&store.Node{})
	orm.Migrator().AutoMigrate(&store.Link{})
	orm.Migrator().AutoMigrate(&store.Object{})

}

func main() {

}

//NewLog Creates creates a new IPFSLog for a given identity
func NewLog(services core_iface.CoreAPI, identity *identityprovider.Identity, options *LogOptions) (*P2PDBLog, error) {
	if services == nil {
		return nil, errmsg.ErrIPFSNotDefined
	}

	if identity == nil {
		return nil, errmsg.ErrIdentityNotDefined
	}

	if options == nil {
		options = &LogOptions{}
	}

	if options.SortFn == nil {
		options.SortFn = clock.LastWriteWins
	}

	maxTime := 0
	if options.Clock != nil {
		maxTime = options.Clock.GetTime()
	}

	maxTime = maxClockTimeForEntries(options.Heads, maxTime)

	// if options.AccessController == nil {
	// 	options.AccessController = &accesscontroller.Default{}
	// }

	// if options.Entries == nil {
	// 	options.Entries = entry.NewOrderedMap()
	// }

	// if len(options.Heads) == 0 && options.Entries.Len() > 0 {
	// 	options.Heads = entry.FindHeads(options.Entries)
	// }

	if options.Concurrency == 0 {
		options.Concurrency = 16
	}

	// if options.IO == nil {
	// 	io, err := cbor.IO(&entry.Entry{}, &entry.LamportClock{})
	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	options.IO = io
	// }

	next := entry.NewOrderedMap()
	for _, key := range options.Entries.Keys() {
		e := options.Entries.UnsafeGet(key)
		for _, n := range e.GetNext() {
			next.Set(n.String(), e)
		}
	}

	return &P2PDBLog{
		Storage:          services,
		ID:               options.ID,
		Identity:         identity,
		AccessController: options.AccessController,
		SortFn:           sorting.NoZeroes(options.SortFn),
		Entries:          options.Entries.Copy(),
		heads:            entry.NewOrderedMapFromEntries(options.Heads),
		Next:             next,
		//	Clock:            clock.NewLamportClock(identity.PublicKey, maxTime),
		io:          options.IO,
		concurrency: options.Concurrency,
	}, nil
}

func (l *P2PDBLog) SetIdentity(identity *identityprovider.Identity) {

}

// maxInt Returns the larger of x or y
func maxInt(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// minInt Returns the larger of x or y
func minInt(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func maxClockTimeForEntries(entries []iface.IPFSLogEntry, defValue int) int {
	max := defValue
	for _, e := range entries {
		max = maxInt(e.GetClock().GetTime(), max)
	}

	return max
}
