package store

import (
	"testing"

	debug "github.com/favframework/debug"
)

func TestDBFactory(t *testing.T) {
	var db *CreateDBFactory
	orm := db.InitDB()

	debug.Dump(orm.Select())
}
