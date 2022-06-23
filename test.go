package main

import (
	"fmt"

	debug "github.com/favframework/debug"
)

type Lister interface {
	Connect() string
	Create() error
	Update() error
	Delete() error
	Select() error
}

type SqliteDB struct{}

// dog实现了Sayer接口

// cat实现了Sayer接口
// func (c cat) lis() {
// 	fmt.Println("喵喵喵")
// }

func (db *SqliteDB) Connect() string {
	return fmt.Sprintf("mongo,地址是")
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

func main() {
	var x Lister    // 声明一个Sayer类型的变量x
	a := SqliteDB{} // 实例化一个cat
	x = &a          // 会报错，原因缺少say方法
	debug.Dump(x.Connect())
}
