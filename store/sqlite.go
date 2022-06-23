package store

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type SqliteDB struct {
	BaseInfo
	OrmDB *gorm.DB
}

func (this *BaseInfo) Init(address string, port int64, account string, passwd string) {
	this.address = address
	this.port = port
	this.account = account
	this.passwd = passwd
}

func (db *SqliteDB) Connect() {
	ormDB, err := gorm.Open("sqlite3", db.BaseInfo.address)
	if err != nil {
		panic("failed to connect database")
	}
	db.OrmDB = ormDB
}

func (db *SqliteDB) Create() error {
	return nil
}
func (db *SqliteDB) Update() error {
	return nil
}
func (db *SqliteDB) Delete() error {
	return nil
}
func (db *SqliteDB) Select() error {
	return nil
}
