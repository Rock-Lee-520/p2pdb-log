package log

import (
	"strconv"
	"sync"
	"time"

	"berty.tech/go-ipfs-log/accesscontroller"
	"berty.tech/go-ipfs-log/entry"
	"berty.tech/go-ipfs-log/entry/sorting"
	"berty.tech/go-ipfs-log/errmsg"
	"berty.tech/go-ipfs-log/iface"
	"berty.tech/go-ipfs-log/io/cbor"
	"github.com/Rock-liyi/p2pdb-log/identityprovider"
	core_iface "github.com/ipfs/interface-go-ipfs-core"
)

type Snapshot = iface.Snapshot
type JSONLog = iface.JSONLog
type LogOptions = iface.LogOptions
type IteratorOptions = iface.IteratorOptions
type IO = iface.IO

type Entry = iface.IPFSLogEntry
type Log = iface.IPFSLog
type AppendOptions = iface.AppendOptions
type SortFn = iface.EntrySortFn

type IPFSLog struct {
	Storage          core_iface.CoreAPI
	ID               string
	AccessController accesscontroller.Interface
	SortFn           iface.EntrySortFn
	Identity         *identityprovider.Identity
	Entries          iface.IPFSLogOrderedEntries
	heads            iface.IPFSLogOrderedEntries
	Next             iface.IPFSLogOrderedEntries
	Clock            iface.IPFSLogLamportClock
	io               iface.IO
	concurrency      uint
	lock             sync.RWMutex
}

func (l *IPFSLog) Len() int {
	return l.Entries.Len()
}

func (l *IPFSLog) RawHeads() iface.IPFSLogOrderedEntries {
	l.lock.RLock()
	heads := l.heads
	l.lock.RUnlock()

	return heads
}

func (l *IPFSLog) IO() IO {
	return l.io
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

// type Log interface {
// 	NewLog(services core_iface.CoreAPI, identity *identityprovider.Identity, options *ipfslog.LogOptions) (*ipfslog.IPFSLog, error)
// 	//Append(ctx context.Context, payload []byte, opts *ipfslog.AppendOptions) (iface.IPFSLogEntry, error)
// }
// type LogFactory struct {
// }

// func (*LogFactory) NewLog(services core_iface.CoreAPI, identity *identityprovider.Identity, options *ipfslog.LogOptions) (*ipfslog.IPFSLog, error) {
// 	return NewLog(services, identity, options)
// }

// func (*LogFactory) Append(ctx context.Context, payload []byte, opts *ipfslog.AppendOptions) (iface.IPFSLogEntry, error) {
// 		ipfslog.IPFSLog.appen
// }

// NewLog Creates creates a new IPFSLog for a given identity
//
// Each IPFSLog gets a unique ID, which can be passed in the options as ID.
//
// Returns a log instance.
//
// ipfs is an instance of IPFS.
//
// identity is an instance of Identity and will be used to sign entries
// Usually this should be a user id or similar.
//
// options.AccessController is an instance of accesscontroller.Interface,
// which by default allows anyone to append to the IPFSLog.
func NewLog(services core_iface.CoreAPI, identity *identityprovider.Identity, options *LogOptions) (*IPFSLog, error) {
	if services == nil {
		return nil, errmsg.ErrIPFSNotDefined
	}

	if identity == nil {
		return nil, errmsg.ErrIdentityNotDefined
	}

	if options == nil {
		options = &LogOptions{}
	}

	if options.ID == "" {
		options.ID = strconv.FormatInt(time.Now().Unix()/1000, 10)
	}

	if options.SortFn == nil {
		options.SortFn = sorting.LastWriteWins
	}

	maxTime := 0
	if options.Clock != nil {
		maxTime = options.Clock.GetTime()
	}
	maxTime = maxClockTimeForEntries(options.Heads, maxTime)

	if options.AccessController == nil {
		options.AccessController = &accesscontroller.Default{}
	}

	if options.Entries == nil {
		options.Entries = entry.NewOrderedMap()
	}

	if len(options.Heads) == 0 && options.Entries.Len() > 0 {
		options.Heads = entry.FindHeads(options.Entries)
	}

	if options.Concurrency == 0 {
		options.Concurrency = 16
	}

	if options.IO == nil {
		io, err := cbor.IO(&entry.Entry{}, &entry.LamportClock{})
		if err != nil {
			return nil, err
		}

		options.IO = io
	}

	next := entry.NewOrderedMap()
	for _, key := range options.Entries.Keys() {
		e := options.Entries.UnsafeGet(key)
		for _, n := range e.GetNext() {
			next.Set(n.String(), e)
		}
	}

	return &IPFSLog{
		Storage:          services,
		ID:               options.ID,
		Identity:         identity,
		AccessController: options.AccessController,
		SortFn:           sorting.NoZeroes(options.SortFn),
		Entries:          options.Entries.Copy(),
		heads:            entry.NewOrderedMapFromEntries(options.Heads),
		Next:             next,
		Clock:            entry.NewLamportClock(identity.PublicKey, maxTime),
		io:               options.IO,
		concurrency:      options.Concurrency,
	}, nil
}
