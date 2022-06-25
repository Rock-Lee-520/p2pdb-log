package core

type Dag interface {
	Compare()
	Merge()
	pull()
	push()
}
