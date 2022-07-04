package core

import (
	"testing"

	"github.com/Rock-liyi/p2pdb-log/encryption"
	debug "github.com/favframework/debug"
	"github.com/stretchr/testify/require"
)

func TestInsertLink(t *testing.T) {
	require := require.New(t)
	var link *LinkFactory
	var cid, err = encryption.GetCid("cid")
	if err != nil {
		debug.Dump(err)
	}

	linkId := cid.String()
	debug.Dump(linkId)
	nodeId := cid.String()
	lastNodeId := cid.String()

	ok, err := link.InsertLink(linkId, nodeId, []string{nodeId, lastNodeId})
	if err != nil {

		debug.Dump(err.Error())
	}
	require.True(ok)
}

func TestLinkDelete(t *testing.T) {
	require := require.New(t)
	var link *LinkFactory
	linkId := "QmU178mY5DBMRvJbU9nZn5nPBr93o7gfDwf5rDJvvohjxq"

	err := link.DeleteLink(linkId)
	if err != nil {
		require.True(false)
	}
	require.True(true)
}
