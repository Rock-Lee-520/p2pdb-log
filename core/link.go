package core

import (
	"github.com/Rock-liyi/p2pdb-log/store"
	debug "github.com/favframework/debug"
)

type Link interface {
	InsertLink(linkId string, nodeId string, lastNodeId string) (bool, error)
	DeleteLink(linkId string) error
}
type LinkFactory struct {
}

/**
 * Insert link into the database
 * creates a new link with the given all of lastNodeIds  and  nodeId
 * returns	 linkId if successful, error otherwise
 **/
func (link *LinkFactory) InsertLink(linkId string, nodeId string, lastNodeIds []string) (bool, error) {
	debug.Dump("====== InsertLink start")

	var linkModel = &store.Link{}

	//Check if the object id is repeated
	debug.Dump(OBJECTID)

	orm := DB.InitDB()
	linkModel.LinkId = linkId
	linkModel.NodeID = nodeId
	linkModel.LastNodeId = "lastNodeId"
	//linkModel.LinkSize = linkSize

	db := orm.Create(&linkModel)
	if err := db.Error; err != nil {
		return false, db.Error
	}

	debug.Dump("====== InsertLink end")
	return true, nil
}

func (link *LinkFactory) DeleteLink(linkId string) error {

	var linkModel = &store.Link{}

	orm := DB.InitDB()
	// linkModel.LinkId = linkId

	err := orm.Where("link_id = ?", linkId).Delete(&linkModel)
	if err != nil {
		return err.Error
	}
	return nil
}
