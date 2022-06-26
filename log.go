package main

type Log interface {
	new(hostID string)
	append()
	join()
	toJson()
	pull()
	push()
}
