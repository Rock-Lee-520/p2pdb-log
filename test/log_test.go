package test

import (
	"context"
	"fmt"
	"testing"

	ipfslog "github.com/Rock-liyi/p2pdb-log"
	idp "github.com/Rock-liyi/p2pdb-log/identityprovider"
	ks "github.com/Rock-liyi/p2pdb-log/keystore"
	"github.com/Rock-liyi/p2pdb-log/log"
	debug "github.com/favframework/debug"
	dssync "github.com/ipfs/go-datastore/sync"
	mocknet "github.com/libp2p/go-libp2p/p2p/net/mock"
	"github.com/stretchr/testify/require"
)

func TestLogCRDT(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	m := mocknet.New(ctx)
	ipfs, closeNode := NewMemoryServices(ctx, t, m)
	defer closeNode()
	//debug.Dump(ipfs)

	datastore := dssync.MutexWrap(NewIdentityDataStore(t))
	keystore, err := ks.NewKeystore(datastore)
	require.NoError(t, err)

	var identities [3]*idp.Identity

	for i, char := range []rune{'A', 'B', 'C'} {
		identity, err := idp.CreateIdentity(ctx, &idp.CreateIdentityOptions{
			Keystore: keystore,
			ID:       fmt.Sprintf("user%c", char),
			Type:     "p2pdb",
		})
		require.NoError(t, err)

		identities[i] = identity
	}
	debug.Dump(identities)

	logA, err := log.NewLog(ipfs, identities[0], &ipfslog.LogOptions{ID: "X"})
	//logB, err := log.NewLog(ipfs, identities[1], &ipfslog.LogOptions{ID: "X"})
	require.NoError(t, err)
	// debug.Dump(logA)
	// debug.Dump(logB)
	str1 := "insert into table(id,name) values(1,'Alice')"
	_, err = logA.Append(ctx, []byte(str1), nil)
	require.NoError(t, err)

	// logA.Join(logB, -1)
	// str2 := "update  table set  name=bob where  id=1"

	// _, err = logA.Append(ctx, []byte(str2), nil)
	// require.NoError(t, err)

	// str3 := "update  table set  name=jack where  id=1"
	// _, err = logB.Append(ctx, []byte(str3), nil)
	// require.NoError(t, err)

	// logB.Join(logA, -1)

	// var data = logB.ToString(nil)
	// fmt.Print(data)

}
