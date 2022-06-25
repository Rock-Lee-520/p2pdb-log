package store

import (
	"time"
)

// gorm.Model definition
type BaseColumn struct {
	CreatedAt time.Time  `gorm:"column:created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at"`
}

// Node model definition
type Node struct {
	BaseColumn
	NodeId             string `gorm:"column:node_id"`
	NodeType           string `gorm:"column:node_type"`
	LamportClock       int64  `gorm:"column:lamport_clock"`
	ReceivingTimestamp int32  `gorm:"column:receiving_timestamp"`
	ReceivingDate      string `gorm:"column:receiving_date"`
	SendingDate        string `gorm:"column:sending_date"`
	SendingTimestamp   int32  `gorm:"column:sending_timestamp"`
	LastNodeId         string `gorm:"column:last_node_id"`
}

//  Object model definition

type Object struct {
	BaseColumn
	ObjectId  string `gorm:"column:object_id"`
	NodeId    string `gorm:"column:node_id"`
	Content   string `gorm:"column:content"`
	Operation string `gorm:"column:operation"`
	Propertie string `gorm:"column:propertie"`
}

type Link struct {
	BaseColumn
	LinkId   string `gorm:"column:link_id"`
	LastCid  string `gorm:"column:last_node_id"`
	NodeID   string `gorm:"column:node_id"`
	LinkSize string `gorm:"column:link_size"`
}
