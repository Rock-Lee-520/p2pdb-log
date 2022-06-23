package store

import (
	"os"
	"path/filepath"

	conf "github.com/Rock-liyi/p2pdb-log/config"
	debug "github.com/favframework/debug"
)

type BaseInfo struct {
	address string
	port    int64
	account string
	passwd  string
}

//Connect()应该返回的是连接对象,这里做了省略,所以用string来代替

type DBconnect interface {
	Init(address string, port int64, account string, passwd string)
	Create() error
	Update() error
	Delete() error
	Select() error
	Connect()
}

type CreateDBFactory struct {
}

func (db *CreateDBFactory) CreateDBConnect(db_type string) DBconnect {
	var orm DBconnect
	switch db_type {
	case "mysql":
		orm = &SqliteDB{}
	default:
		orm = &SqliteDB{}
	}
	return orm
}

func (db *CreateDBFactory) InitDB() DBconnect {
	var connect = db.CreateDBConnect("sqlite")
	name := "p2pdb_log"
	//init config,get db path
	dataPath := conf.GetDataPath()
	if dataPath != "" {
		dataPath = dataPath + "/"
	}

	binary, _ := os.Getwd()
	root := filepath.Dir(binary)
	if root != "" && dataPath == "" {
		dataPath = root + "/"
	}
	debug.Dump(dataPath + name + ".db")
	address := dataPath + name + ".db"
	connect.Init(address, 0, "", "")
	connect.Connect()
	return connect

}
