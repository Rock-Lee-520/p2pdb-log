package core

type Link interface {
	InsertLink(node_id string, node_type string, lamport_clock int64, receiving_timestamp int32,
		receiving_date string, sending_date string, send_timestamp int32, last_id string) (bool, error)
}

type LinkFactory interface {
}

func init() {
	//	orm = DB.InitDB()
	//	debug.Dump("call link init function=====")
}

// func (link *LinkFactory) InsertLink(node_id string, node_type string, lamport_clock int64, receiving_timestamp int32,
// 	receiving_date string, sending_date string, send_timestamp int32, last_id string) (bool, error) {
// 	debug.Dump(link)
// 	return 1, nil
// }
