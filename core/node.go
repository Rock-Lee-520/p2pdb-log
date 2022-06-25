package core

import (
	"github.com/Rock-liyi/p2pdb-log/store"
	debug "github.com/favframework/debug"
)

var DB *store.CreateDBFactory
var orm store.DBconnect

type Node interface {
	InsertNode(node_id string, node_type string, lamport_clock int64, receiving_timestamp int32,
		receiving_date string, sending_date string, send_timestamp int32, last_id string) (int, error)
	FindNode()
	DeleteNode()
	UpdateNode()
	// Init(address string, port int64, account string, passwd string)
	// Create() error
	// Update() error
	// Delete() error
	// Select(query interface{}, args ...interface{}) *gorm.DB
	// Where(query interface{}, args ...interface{}) *gorm.DB
	// First(out interface{}, where ...interface{}) *gorm.DB
	// Find(out interface{}, where ...interface{}) *gorm.DB
	// Connect()
}

type NodeFactory struct {
}

func init() {
	orm = DB.InitDB()
	debug.Dump("call init function=====")
}

// func (nodeFactory *NodeFactory) init() {

// 	nodeFactory.init()
// }

func (node *NodeFactory) InsertNode(node_id string, node_type string, lamport_clock int64, receiving_timestamp int32,
	receiving_date string, sending_date string, send_timestamp int32, last_id string) (int, error) {
	var nodeModel = &store.Node{}
	nodeModel.NodeId = node_id
	nodeModel.NodeType = node_type
	nodeModel.LamportClock = lamport_clock
	nodeModel.ReceivingTimestamp = receiving_timestamp
	nodeModel.ReceivingDate = receiving_date
	nodeModel.SendingDate = sending_date
	nodeModel.SendingTimestamp = send_timestamp
	nodeModel.LastId = last_id

	var result = orm.Create(&nodeModel)
	return 0, nil
}
