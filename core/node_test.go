package core

import (
	"testing"

	debug "github.com/favframework/debug"
)

func TestDBSelect(t *testing.T) {

	var node *NodeFactory
	var data = node.InsertNode()

	debug.Dump(data)
}
