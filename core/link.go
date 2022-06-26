package core

import (
	"github.com/Rock-liyi/p2pdb-log/store"
	debug "github.com/favframework/debug"
)

type Link interface {
	InsertLink(linkId string, nodeId string, lastNodeId string) (bool, error)
}
type LinkFactory struct {
}

func (link *LinkFactory) InsertLink(linkId string, nodeId string, lastNodeId string) (bool, error) {
	debug.Dump("====== InsertLink start")

	var linkModel = &store.Link{}

	//Check if the object id is repeated
	debug.Dump(OBJECTID)

	orm := DB.InitDB()
	linkModel.LinkId = linkId
	linkModel.NodeID = nodeId
	linkModel.LastNodeId = lastNodeId
	//linkModel.LinkSize = linkSize

	db := orm.Create(&linkModel)
	if err := db.Error; err != nil {
		return false, db.Error
	}

	debug.Dump("====== InsertLink end")
	return true, nil
}
