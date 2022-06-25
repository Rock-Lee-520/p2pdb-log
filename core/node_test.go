package core

import (
	"testing"

	"github.com/Rock-liyi/p2pdb-log/encryption"
	debug "github.com/favframework/debug"
	"github.com/stretchr/testify/require"
)

func TestDBSelect(t *testing.T) {
	require := require.New(t)
	var node *NodeFactory
	var cid, err = encryption.GetCid("cid")
	if err != nil {
		debug.Dump(err)
	}

	nodeId := cid.String()
	nodeType := "log"
	lamportClock := 123
	receivingTimestamp := 123
	receivingDate := "2021"
	sendingDate := "2021"
	sendingTimestamp := 123
	lastId := cid.String()

	ok, err := node.InsertNode(nodeId, nodeType, int64(lamportClock), int32(receivingTimestamp),
		receivingDate, sendingDate, int32(sendingTimestamp), lastId)
	if err != nil {
		debug.Dump(err.Error())
	}
	require.True(ok)
}
