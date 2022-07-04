package core

import (
	"testing"

	"github.com/Rock-liyi/p2pdb-log/encryption"
	debug "github.com/favframework/debug"
	"github.com/stretchr/testify/require"
)

func TestObjectInsert(t *testing.T) {
	require := require.New(t)
	var object *ObjectFactory
	var cid, err = encryption.GetCid("cid")
	if err != nil {
		debug.Dump(err)
	}

	objectId := cid.String()
	nodeId := cid.String()
	content := "content"
	operation := "operation"
	property := "property"
	ok, err := object.InsertObject(objectId, nodeId, content,
		operation, property)
	if err != nil {
		debug.Dump("======")
		debug.Dump(err.Error())
	}
	require.True(ok)
}
