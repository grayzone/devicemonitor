package conn

import (
	"time"

	"log"

	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

type MessageType int64

const (
	REQUEST  MessageType = 0
	RESPONSE MessageType = 1
)

type MessageStatus int64

const (
	NONE      MessageStatus = 0
	PROCESSED MessageStatus = 1
	INVALID   MessageStatus = 2
)

const ()

type Message struct {
	Id          int64
	Messagetype MessageType
	Info        string
	Status      MessageStatus
	Updatetime  time.Time
}

type DeviceID int64

const (
	ForceTraid        DeviceID = 0x00
	ValleylabExchange DeviceID = 0x01
	PatriotGenerator  DeviceID = 0x02
	IntegratedOR      DeviceID = 0x3D
	ServiceApps       DeviceID = 0xD8
)

type Setting struct {
	Id             int64
	Isconnected    bool
	Deviceid       DeviceID
	Protocolver    string
	Sessionkey     string
	Sequence       string
	Writeinterval  int
	Sessionstatus  uint32
	Sessiontimeout uint32
	Messagetimeout uint32
	Maxretrycount  uint32
	Updatetime     time.Time
}

func init() {
	orm.Debug = false
	orm.RegisterDriver("postgres", orm.DR_Postgres)
	connstr := "user=postgres password=123456 dbname=devicemonitor sslmode=disable"
	orm.RegisterDataBase("default", "postgres", connstr)

	orm.RegisterModel(new(Message), new(Setting))

}

func (m *Message) InsertMessage() {
	o := orm.NewOrm()
	o.Begin()

	id, err := o.Insert(m)
	if err != nil {
		log.Println(err.Error())
		o.Rollback()
	} else {
		//		log.Println(id)
		m.Id = id
	}

	o.Commit()
}

func (m *Message) Get() error {
	o := orm.NewOrm()
	err := o.Read(m)
	return err

}

func (m *Message) GetOneRequest() error {
	o := orm.NewOrm()

	err := o.QueryTable("message").Filter("messagetype", REQUEST).Filter("status", NONE).OrderBy("id").Limit(1).One(m)

	return err
}

func (m *Message) GetOneResponse() error {
	o := orm.NewOrm()
	err := o.QueryTable("message").Filter("messagetype", RESPONSE).Filter("status", NONE).OrderBy("id").Limit(1).One(m)

	return err
}

func (m *Message) UpdateStatus() error {
	o := orm.NewOrm()
	_, err := o.Update(m, "Status")
	return err
}

func (m *Message) UpdateInfo() error {
	o := orm.NewOrm()
	_, err := o.Update(m, "Info")
	return err
}

func (m *Message) DeleteMessage() {

	o := orm.NewOrm()

	_, err := o.Delete(m)
	if err != nil {
		log.Println("DeleteMessage:", err.Error())
	}
}

func (s *Setting) GetSetting() error {
	o := orm.NewOrm()
	err := o.QueryTable("setting").Limit(1).One(s)
	return err
}

func (s *Setting) Update(item string) error {
	o := orm.NewOrm()
	_, err := o.QueryTable("setting").Update(orm.Params{
		"isconnected":    s.Isconnected,
		"deviceid":       s.Deviceid,
		"protocolver":    s.Protocolver,
		"sessionkey":     s.Sessionkey,
		"sequence":       s.Sequence,
		"writeinterval":  s.Writeinterval,
		"sessionstatus":  s.Sessionstatus,
		"sessiontimeout": s.Sessiontimeout,
		"messagetimeout": s.Messagetimeout,
		"maxretrycount":  s.Maxretrycount,
		"updatetime":     time.Now(),
	})
	return err
}

func (s *Setting) UpdateIsConnected() error {
	o := orm.NewOrm()
	_, err := o.QueryTable("setting").Update(orm.Params{
		"isconnected": s.Isconnected,
		"updatetime":  time.Now(),
	})
	return err
}

func (s *Setting) UpdateDeviceid() error {
	o := orm.NewOrm()
	_, err := o.QueryTable("setting").Update(orm.Params{

		"deviceid":   s.Deviceid,
		"updatetime": time.Now(),
	})
	return err
}

func (s *Setting) UpdateProtocolVer() error {
	o := orm.NewOrm()
	_, err := o.QueryTable("setting").Update(orm.Params{

		"protocolver": s.Protocolver,
		"updatetime":  time.Now(),
	})
	return err
}

func (s *Setting) UpdateSessionKey() error {
	o := orm.NewOrm()
	_, err := o.QueryTable("setting").Update(orm.Params{

		"sessionkey": s.Sessionkey,
		"updatetime": time.Now(),
	})
	return err
}

func (s *Setting) UpdateSequence() error {
	o := orm.NewOrm()
	_, err := o.QueryTable("setting").Update(orm.Params{

		"sequence":   s.Sequence,
		"updatetime": time.Now(),
	})
	return err
}

func (s *Setting) UpdateWriteInterval() error {
	o := orm.NewOrm()
	_, err := o.QueryTable("setting").Update(orm.Params{

		"writeinterval": s.Writeinterval,
		"updatetime":    time.Now(),
	})
	return err
}

func (s *Setting) UpdateSessionStatus() error {
	o := orm.NewOrm()
	_, err := o.QueryTable("setting").Update(orm.Params{

		"sessionstatus": s.Sessionstatus,
		"updatetime":    time.Now(),
	})
	return err

}

func (s *Setting) UpdateSessiontimeout() error {
	o := orm.NewOrm()
	_, err := o.QueryTable("setting").Update(orm.Params{

		"sessiontimeout": s.Sessiontimeout,
		"updatetime":     time.Now(),
	})
	return err
}

func (s *Setting) UpdateMessagetimeout() error {
	o := orm.NewOrm()
	_, err := o.QueryTable("setting").Update(orm.Params{

		"messagetimeout": s.Messagetimeout,
		"updatetime":     time.Now(),
	})
	return err
}

func (s *Setting) UpdateMaxretrycount() error {
	o := orm.NewOrm()
	_, err := o.QueryTable("setting").Update(orm.Params{

		"maxretrycount": s.Maxretrycount,
		"updatetime":    time.Now(),
	})
	return err
}
