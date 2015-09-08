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
)

const ()

type Message struct {
	Id          int
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
	Id            int
	Isconnected   bool
	Deviceid      DeviceID
	Protocolver   string
	Sessionkey    string
	Sequence      string
	Writeinterval int64
}

func init() {
	orm.Debug = true
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
		log.Println(id)
	}

	o.Commit()
}

func (m *Message) GetOneRequest() error {
	o := orm.NewOrm()

	err := o.QueryTable("message").Filter("messagetype", REQUEST).Filter("status", NONE).OrderBy("id").Limit(1).One(m)
	return err

}

func (m *Message) DeleteMessage() {

	o := orm.NewOrm()

	_, err := o.Delete(m)
	if err != nil {
		log.Fatal(err)
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
		"isconnected":   s.Isconnected,
		"deviceid":      s.Deviceid,
		"Protocolver":   s.Protocolver,
		"Sessionkey":    s.Sessionkey,
		"Sequence":      s.Sequence,
		"Writeinterval": s.Writeinterval,
	})
	return err
}

func (s *Setting) UpdateIsConnected() error {
	o := orm.NewOrm()
	_, err := o.QueryTable("setting").Update(orm.Params{
		"isconnected": s.Isconnected,
	})
	return err
}

func (s *Setting) UpdateDeviceid() error {
	o := orm.NewOrm()
	_, err := o.QueryTable("setting").Update(orm.Params{

		"deviceid": s.Deviceid,
	})
	return err
}

func (s *Setting) UpdateProtocolVer() error {
	o := orm.NewOrm()
	_, err := o.QueryTable("setting").Update(orm.Params{

		"Protocolver": s.Protocolver,
	})
	return err
}

func (s *Setting) UpdateSessionKey() error {
	o := orm.NewOrm()
	_, err := o.QueryTable("setting").Update(orm.Params{

		"Sessionkey": s.Sessionkey,
	})
	return err
}

func (s *Setting) UpdateSequence() error {
	o := orm.NewOrm()
	_, err := o.QueryTable("setting").Update(orm.Params{

		"Sequence": s.Sequence,
	})
	return err
}

func (s *Setting) UpdateWriteInterval() error {
	o := orm.NewOrm()
	_, err := o.QueryTable("setting").Update(orm.Params{

		"Writeinterval": s.Writeinterval,
	})
	return err
}
