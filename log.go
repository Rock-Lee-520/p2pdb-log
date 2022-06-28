package main

import "p2pdb-log/store"

type Log interface {
	new(hostID string)
	append()
	join()
	toJson()
	pull()
	push()
}

func init() {

	var db *store.CreateDBFactory
	orm := db.InitDB()
	// init node、link、object table schema
	orm.Migrator().AutoMigrate(&store.Node{})
	orm.Migrator().AutoMigrate(&store.Link{})
	orm.Migrator().AutoMigrate(&store.Object{})

}

func main() {

}
