package core

import (
	ipfslog "berty.tech/go-ipfs-log"
	"berty.tech/go-ipfs-log/identityprovider"
	core_iface "github.com/ipfs/interface-go-ipfs-core"
)

type Log interface {
	NewLog(services core_iface.CoreAPI, identity *identityprovider.Identity, options *ipfslog.LogOptions) (*ipfslog.IPFSLog, error)
	//Append(ctx context.Context, payload []byte, opts *ipfslog.AppendOptions) (iface.IPFSLogEntry, error)
}
type LogFactory struct {
}

func (*LogFactory) NewLog(services core_iface.CoreAPI, identity *identityprovider.Identity, options *ipfslog.LogOptions) (*ipfslog.IPFSLog, error) {
	return ipfslog.NewLog(services, identity, options)
}

// func (*LogFactory) Append(ctx context.Context, payload []byte, opts *ipfslog.AppendOptions) (iface.IPFSLogEntry, error) {
// 		ipfslog.IPFSLog.appen
// }
