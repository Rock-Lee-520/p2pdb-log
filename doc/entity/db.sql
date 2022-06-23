-- node definition

CREATE TABLE node (
	node_id TEXT NOT NULL,
	"type" TEXT,
	lamport_clock TEXT,
	receiving_timestamp INTEGER,
	receiving_date TEXT,
	sending_date TEXT,
	send_timestamp INTEGER,
	last_id TEXT
);

