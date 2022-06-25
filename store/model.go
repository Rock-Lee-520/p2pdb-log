package store

import (
	"time"
)

// gorm.Model definition
type Model struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

// Node model definition
type Node struct {
	NodeId             string `gorm:"column:node_id"`
	NodeType           string `gorm:"column:node_type"`
	LamportClock       int64  `gorm:"column:lamport_clock"`
	ReceivingTimestamp int32  `gorm:"column:receiving_timestamp"`
	ReceivingDate      string `gorm:"column:receiving_date"`
	SendingDate        string `gorm:"column:sending_date"`
	SendingTimestamp   int32  `gorm:"column:sending_timestamp"`
	LastId             string `gorm:"column:last_id"`
}

//  Object model definition
type Object struct {
	ObjectId         string    `gorm:"column:object_id"`
	NodeId           string    `gorm:"column:node_id"`
	CreatedTimestamp time.Time `gorm:"column:created_timestamp"`
	CreatedDate      string    `gorm:"column:created_date"`
	Operation        string    `gorm:"column:operation"`
	Propertie        string    `gorm:"column:propertie"`
}

type Link struct {
	LinkId           string    `gorm:"column:link_id"`
	LastCid          string    `gorm:"column:last_node_id"`
	NodeID           string    `gorm:"column:node_id"`
	LinkSize         string    `gorm:"column:link_size"`
	CreatedTimestamp time.Time `gorm:"column:created_timestamp"`
	CreatedDate      string    `gorm:"column:created_date"`
}
