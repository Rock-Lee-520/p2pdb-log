-- node definition

CREATE TABLE "node" (
	node_id TEXT NOT NULL,
	node_type TEXT,
	lamport_clock INTEGER,
	receiving_timestamp INTEGER,
	receiving_date TEXT,
	sending_date TEXT,
	send_timestamp INTEGER,
	last_node_id TEXT
);

-- link definition

CREATE TABLE "link" (
	link_id TEXT,
	last_node_id TEXT,
	node_id TEXT,
	link_size INTEGER,
	created_timestamp TEXT,
	created_date TEXT
);



-- "object" definition

CREATE TABLE "object" (
	object_id TEXT NOT NULL,
	node_id TEXT NOT NULL,
	created_timestamp NUMERIC,
	created_date TEXT,
	operation TEXT,
	propertie TEXT
);