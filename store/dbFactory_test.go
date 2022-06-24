package store

import (
	"testing"

	debug "github.com/favframework/debug"
)

func TestDBFactory(t *testing.T) {

	var db *CreateDBFactory
	orm := db.InitDB()
	var node = &Node{}
	orm.Where("name = ?", "jinzhu").First(&node)

	debug.Dump(node)
}
