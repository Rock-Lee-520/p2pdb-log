package src

import (
	"github.com/Rock-liyi/p2pdb-log/store"
)

var DB *store.CreateDBFactory

type Node interface {
	InsertNode(node_id string, node_type string, lamport_clock int64, receiving_timestamp int32,
		receiving_date string, sending_date string, sending_timestamp int32, last_node_id string) (bool, error)
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

func (node *NodeFactory) InsertNode(nodeId string, nodeType string, lamportClock int64, receivingTimestamp int32,
	receivingDate string, sendingDate string, sendingTimestamp int32, lastNodeId string) (bool, error) {
	var nodeModel = &store.Node{}
	nodeModel.NodeId = nodeId
	nodeModel.NodeType = nodeType
	nodeModel.LamportClock = lamportClock
	nodeModel.ReceivingTimestamp = receivingTimestamp
	nodeModel.ReceivingDate = receivingDate
	nodeModel.SendingDate = sendingDate
	nodeModel.SendingTimestamp = sendingTimestamp
	nodeModel.LastNodeId = lastNodeId
	orm := DB.InitDB()
	err := orm.Create(&nodeModel)
	if err.Error != nil {
		return false, err.Error
	}
	return true, nil
}
