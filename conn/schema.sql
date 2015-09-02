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