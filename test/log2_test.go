package test

import (
	"p2pdb-log/store"

	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestInit(t *testing.T) {
	type Product struct {
		gorm.Model
		Code  string
		Price uint
	}

	DB, err := gorm.Open(sqlite.Open("/Users/rockli/go/src/p2pdb-server/data/test.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	var bool = DB.Migrator().HasTable(&Product{})

	if bool != true {
		DB.Migrator().CreateTable(&Product{})
	}

}

func TestTable(t *testing.T) {
	type Product struct {
		gorm.Model
		Code  string
		Price uint
	}
	var db *store.CreateDBFactory
	orm := db.InitDB()

	orm.Migrator().AutoMigrate(&Product{})
}
