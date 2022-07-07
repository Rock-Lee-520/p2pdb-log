package src

import (
	"github.com/Rock-liyi/p2pdb-log/store"
	debug "github.com/favframework/debug"
)

type Object interface {

	// it need  to  note that  if the object content is repeated, only one can be saved,
	// and  then return the repeatedly object id
	InsertObject(orm store.DBconnect, objectId string, nodeId string, Content string, Operation string,
		Propertie string) (bool, error)
}

const (
	OBJECTID = "object_id"
)

type ObjectFactory struct {
}

func (object *ObjectFactory) InsertObject(objectId string, nodeId string, content string, operation string,
	property string) (bool, error) {
	debug.Dump("====== InsertObject start")
	var objectModel = &store.Object{}

	//Check if the object id is repeated
	debug.Dump(OBJECTID)

	orm := DB.InitDB()
	objectModel.NodeId = nodeId
	objectModel.ObjectId = objectId
	objectModel.Operation = operation
	objectModel.Content = content
	objectModel.Property = property
	db := orm.Where(OBJECTID+" = ?", objectId).FirstOrCreate(&objectModel)
	//debug.Dump(objectModel)
	if err := db.Error; err != nil {
		return false, db.Error
	}

	debug.Dump("====== InsertObject end")
	return true, nil
}
