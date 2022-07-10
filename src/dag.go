package src

type Dag interface {
	Compare()
	Merge()
	pull()
	push()
}
