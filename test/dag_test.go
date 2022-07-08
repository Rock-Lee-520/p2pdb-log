package test

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/Rock-liyi/p2pdb-log/entry/sorting"
	idp "github.com/Rock-liyi/p2pdb-log/identityprovider"
	"github.com/Rock-liyi/p2pdb-log/iface"
	ks "github.com/Rock-liyi/p2pdb-log/keystore"
	"github.com/Rock-liyi/p2pdb-log/log"
	ipfslog "github.com/Rock-liyi/p2pdb-log/log"
	debug "github.com/favframework/debug"
	dssync "github.com/ipfs/go-datastore/sync"
	mocknet "github.com/libp2p/go-libp2p/p2p/net/mock"
	"github.com/stretchr/testify/require"
)

func TestDagUtils(t *testing.T) {
	require := require.New(t)
	debug.Dump(strconv.FormatInt(time.Now().Unix()/1000, 10))
	require.True(true)
}

func TestDag(t *testing.T) {
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
			//ID:   "A",
			Type: "p2pdb",
		})
		require.NoError(t, err)
		//	debug.Dump(char)
		identities[i] = identity
	}
	// next := []cid.Cid{}

	logA, err := log.NewLog(ipfs, identities[0], &ipfslog.LogOptions{ID: "A"})
	logB, err := log.NewLog(ipfs, identities[1], &ipfslog.LogOptions{ID: "A"})
	require.NoError(t, err)

	debug.Dump("logA.GetTime() is ")
	debug.Dump(logA.Clock.GetTime())

	str1 := "A1"
	_, err = logA.AppendByNewTime(ctx, []byte(str1), nil, 1)
	require.NoError(t, err)

	str2 := "A2"
	_, err = logA.AppendByNewTime(ctx, []byte(str2), nil, 2)
	require.NoError(t, err)

	str3 := "A3"
	_, err = logA.AppendByNewTime(ctx, []byte(str3), nil, 3)
	require.NoError(t, err)

	str4 := "A4"
	_, err = logA.AppendByNewTime(ctx, []byte(str4), nil, 4)
	require.NoError(t, err)

	str5 := "A5"
	_, err = logA.AppendByNewTime(ctx, []byte(str5), nil, 5)
	require.NoError(t, err)
	debug.Dump("logA.GetTime() is ")
	debug.Dump(logA.Clock.GetTime())

	strB1 := "B1"
	_, err = logB.AppendByNewTime(ctx, []byte(strB1), nil, 1)
	require.NoError(t, err)

	strB2 := "B2"
	_, err = logB.AppendByNewTime(ctx, []byte(strB2), nil, 2)
	require.NoError(t, err)

	strB3 := "B3"
	_, err = logB.AppendByNewTime(ctx, []byte(strB3), nil, 3)
	require.NoError(t, err)

	strB4 := "B4"
	_, err = logB.AppendByNewTime(ctx, []byte(strB4), nil, 4)
	require.NoError(t, err)

	strB5 := "B5"
	_, err = logB.AppendByNewTime(ctx, []byte(strB5), nil, 5)
	require.NoError(t, err)
	debug.Dump("logB.GetTime() is ")
	debug.Dump(logB.Clock.GetTime())

	var dataA = logA.ToString(nil)
	var dataB = logB.ToString(nil)
	fmt.Print(dataA)
	fmt.Print("\r\n ============\r\n ")
	fmt.Print(dataB)

}

func TestDagMerge(t *testing.T) {
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
			//ID:   "A",
			Type: "p2pdb",
		})
		require.NoError(t, err)
		//	debug.Dump(char)
		identities[i] = identity
	}
	// next := []cid.Cid{}

	logA, err := log.NewLog(ipfs, identities[0], &ipfslog.LogOptions{ID: "A", SortFn: sorting.Compare})
	//logB, err := log.NewLog(ipfs, identities[0], &ipfslog.LogOptions{ID: "A"})
	require.NoError(t, err)
	logA.Clock.SetTime(80)
	debug.Dump("logA.GetTime() is ")
	debug.Dump(logA.Clock.GetTime())

	str1 := "A1"
	_, err = logA.AppendByNewTime(ctx, []byte(str1), nil, 80)
	require.NoError(t, err)
	str2 := "A2"
	_, err = logA.AppendByNewTime(ctx, []byte(str2), nil, 78)
	require.NoError(t, err)
	str3 := "A3"
	_, err = logA.AppendByNewTime(ctx, []byte(str3), nil, 79)

	require.NoError(t, err)
	debug.Dump("logA.GetTime() is ")
	debug.Dump(logA.Clock.GetTime())

	// var opts *iface.AppendOptions
	// opts.PointerCount = 77
	logA.Clock.SetTime(60)
	strB1 := "B1"
	_, err = logA.AppendByNewTime(ctx, []byte(strB1), &iface.AppendOptions{PointerCount: 60}, 60)
	require.NoError(t, err)

	strB2 := "B2"
	_, err = logA.AppendByNewTime(ctx, []byte(strB2), nil, 62)
	require.NoError(t, err)

	strB3 := "B3"
	_, err = logA.AppendByNewTime(ctx, []byte(strB3), nil, 62)
	require.NoError(t, err)

	require.NoError(t, err)
	debug.Dump("logA.GetTime() is ")
	debug.Dump(logA.Clock.GetTime())
	//logA.Join(logB, -1)
	//logB.Join(logA, -1)
	var dataA = logA.ToString(nil)
	//	var dataB = logB.ToString(nil)

	fmt.Print(dataA)
	fmt.Print("\r\n ============\r\n ")
	//fmt.Print(dataB)

}
