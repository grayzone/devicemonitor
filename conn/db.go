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
	Id          int64
	Messagetype MessageType
	Info        string
	Status      MessageStatus
	Updatetime  time.Time
}

func init() {
	orm.Debug = true
	orm.RegisterDriver("postgres", orm.DR_Postgres)
	connstr := "user=postgres password=123456 dbname=devicemonitor sslmode=disable"
	orm.RegisterDataBase("default", "postgres", connstr)

	orm.RegisterModel(new(Message))

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
