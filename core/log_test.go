package core

import (
	"context"
	"testing"

	debug "github.com/favframework/debug"
	_ "github.com/ipld/go-ipld-prime/codec/dagcbor"
	mocknet "github.com/libp2p/go-libp2p/p2p/net/mock"
)

func TestLogCRDT(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	m := mocknet.New(ctx)
	// ipfs, closeNode := NewMemoryServices(ctx, t, m)
	// defer closeNode()
	debug.Dump(m)
	// datastore := dssync.MutexWrap(NewIdentityDataStore(t))
	// keystore, err := ks.NewKeystore(datastore)
	// require.NoError(t, err)

	// var identities [3]*idp.Identity

}
