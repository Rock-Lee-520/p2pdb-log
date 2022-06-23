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
	NodeId             string    `gorm:"column:node_id"`
	Type               string    `gorm:"column:type"`
	LamportClock       string    `gorm:"column:lamport_clock"`
	ReceivingTimestamp time.Time `gorm:"column:receiving_timestamp"`
	ReceivingDate      string    `gorm:"column:receiving_date"`
	SendingDate        string    `gorm:"column:sending_date"`
	SendingTimestamp   time.Time `gorm:"column:sending_timestamp"`
	LastId             string    `gorm:"column:last_id"`
}

//  Object model definition
type Object struct {
	ObjectId         string    `gorm:"column:object_id"`
	Cid              string    `gorm:"column:cid"`
	CreatedTimestamp time.Time `gorm:"column:created_timestamp"`
	CreatedDate      string    `gorm:"column:created_date"`
	Operation        string    `gorm:"column:operation"`
	Propertie        string    `gorm:"column:propertie"`
}

type Link struct {
	LinkId           string    `gorm:"column:link_id"`
	LastCid          string    `gorm:"column:last_cid"`
	Cid              string    `gorm:"column:cid"`
	Size             string    `gorm:"column:size"`
	createdTimestamp time.Time `gorm:"column:created_timestamp"`
	createdDate      string    `gorm:"column:created_date"`
}
