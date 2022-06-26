package core

import (
	"testing"

	"github.com/Rock-liyi/p2pdb-log/encryption"
	debug "github.com/favframework/debug"
	"github.com/stretchr/testify/require"
)

func TestLinkInsert(t *testing.T) {
	require := require.New(t)
	var link *LinkFactory
	var cid, err = encryption.GetCid("cid")
	if err != nil {
		debug.Dump(err)
	}

	linkId := cid.String()
	nodeId := cid.String()
	lastNodeId := cid.String()
	ok, err := link.InsertLink(linkId, nodeId, lastNodeId)
	if err != nil {
		debug.Dump("======")
		debug.Dump(err.Error())
	}
	require.True(ok)
}
