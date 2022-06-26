package core

import (
	"testing"

	debug "github.com/favframework/debug"
)

func TestGetPublick(t *testing.T) {
	//require := require.New(t)
	var identity *Identity
	var key, _ = identity.GetPublicKey()
	debug.Dump(key)
	//require.True(ok)
}
