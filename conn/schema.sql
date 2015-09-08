DROP table message;
CREATE table IF NOT EXISTS message (
  id serial NOT NULL,
  messagetype integer,
  info text,
  status integer,
  updatetime timestamp with time zone,
  createtime timestamp with time zone default now(),
  CONSTRAINT message_pkey PRIMARY KEY (id)
)WITH (
  OIDS=FALSE
);

DROP table setting;
CREATE table IF NOT EXISTS setting(
	id serial NOT NULL,
	isconnected boolean,
	deviceid integer,
	protocolver text,
	sessionkey text,
	sequence text,
	writeinterval integer,
	CONSTRAINT setting_pkey PRIMARY KEY (id)
)WITH(
	OIDS = FALSE
);
INSERT INTO setting(
            isconnected, deviceid, protocolver, sessionkey, sequence, writeinterval)
    VALUES (FALSE, 1, '1.0025', 'FF', '0', 200);