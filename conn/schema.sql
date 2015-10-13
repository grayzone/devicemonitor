DROP table IF EXISTS message;
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

DROP table IF EXISTS setting;
CREATE table IF NOT EXISTS setting(
	id serial NOT NULL,
	isconnected boolean,
	deviceid integer,
	protocolver text,
	sessionkey text,
	sequence text,
	writeinterval integer,
  sessionstatus  integer,
  sessiontimeout integer,
  messagetimeout integer,
  maxretrycount  integer,
  devicename text,
  sensorbroadcastperiod integer,
  updatetime timestamp with time zone,
	CONSTRAINT setting_pkey PRIMARY KEY (id)
)WITH(
	OIDS = FALSE
);
INSERT INTO setting(
            isconnected, deviceid, protocolver, sessionkey, sequence, writeinterval, sessionstatus,sessiontimeout,messagetimeout,maxretrycount)
    VALUES (FALSE, 1, '1.0025', 'FF', '0', 200, 0,0,0,0);


DROP table IF EXISTS sensordata;
CREATE table IF NOT EXISTS sensordata(
  id serial NOT NULL,
  isvaliddata boolean,
  sequencenumber bigint,
  isactivatingflag boolean,
  vavg real,
  iavg real,
  pavg real,
  vrms real,
  irms real,
  viphase real,
  vpk real,
  ipk real,
  vcf real,
  icf real,
  zload real,
  t1 real,
  t2 real,
  leakage real,
  stimpos real,
  stimneg real,
  oltarget real,
  createtime timestamp with time zone default now(),
  CONSTRAINT sensordata_pkey PRIMARY KEY (id)

) WITH(
  OIDS = FALSE
);