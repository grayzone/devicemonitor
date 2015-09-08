package main

import (
	"encoding/hex"
	"log"
	"sync"
	"time"

	"github.com/grayzone/devicemonitor/comm"
	"github.com/grayzone/devicemonitor/conn"
	"github.com/grayzone/devicemonitor/ftprotocol"
)

func IncreaseSeq(seq byte) byte {
	if seq == 0x39 {
		return 0x30
	}
	return seq + 1
}

func Init() {

}

func Test_GetSensorData() {
	//	var serial comm.Serial
	err := comm.OpenSerial()

	if err != nil {
		log.Println(err.Error())
	}
	var sessionkey []byte = []byte{0x46, 0x46}
	var sequence byte = 0x30
	for i := 0; ; i++ {

		if i == 0 {
			var r ftprotocol.RequestSession
			r.DeviceID = 0xD8
			r.ProtocolVersion = 0x2729
			r.SessionKey = sessionkey
			r.Sequence = sequence

			r.NoAck = true
			log.Printf("send requestsession : %X", r.Message())

			comm.Writer(r.Message())

			res, _ := comm.Reader()
			log.Printf("received requestsession:%X", res)
			if len(res) > 13 {
				var s ftprotocol.RequestSessionResponse
				s.Parse(res[2:])
				sessionkey = s.SessionKey
				sequence = s.Sequence
				log.Println("session status:", s.SessionStatus)
				log.Println("device id:", s.DeviceID)
				log.Println("protocol version:", s.ProtocolVersion)
			}

		}
		var k ftprotocol.KeepAlive
		k.SessionKey = sessionkey
		k.Sequence = sequence
		log.Printf("send keepalive : %X", k.Message())
		comm.Writer(k.Message())

		res, _ := comm.Reader()
		log.Printf("received keepalive:%X", res)
		if len(res) > 1 {
			if res[0] == 0x06 {
				sequence = res[1]
			}
		}
		sequence = IncreaseSeq(sequence)

		if i == 10 {
			var s ftprotocol.GetSensor
			s.SessionKey = sessionkey
			s.Sequence = sequence
			s.IsBroadcast = true
			s.Broadcastperiod = 10
			s.IsAllSensorData = true
			log.Printf("send GetSensor : %X", s.Message())
			comm.Writer(s.Message())

			res, _ := comm.Reader()
			log.Printf("received GetSensor:%X", res)
			sequence = IncreaseSeq(sequence)
		}

		if i == 20 {
			var s ftprotocol.GetSensor
			s.SessionKey = sessionkey
			s.Sequence = sequence
			s.IsBroadcast = false
			s.Broadcastperiod = 10
			s.IsAllSensorData = true
			log.Printf("send GetSensor : %X", s.Message())
			comm.Writer(s.Message())

			res, _ := comm.Reader()
			log.Printf("received GetSensor:%X", res)
			sequence = IncreaseSeq(sequence)

		}

	}

}

func generator(t time.Duration) {
	var sequence byte = 0x30
	for {
		var s conn.Setting
		s.GetSetting()
		if s.Isconnected {
			var k ftprotocol.KeepAlive
			k.SessionKey = []byte(s.Sessionkey)
			k.Sequence = byte(s.Sequence[0])

			log.Printf("send keepalive : %X", k.Message())

			var m conn.Message
			m.Messagetype = conn.REQUEST
			m.Info = hex.EncodeToString(k.Message())
			m.Status = conn.NONE

			m.InsertMessage()

			sequence = IncreaseSeq(sequence)

		}

		time.Sleep(time.Millisecond * time.Duration(s.Writeinterval))

	}

}

func writer(t time.Duration) {
	// get one request
	for {
		var m conn.Message
		err := m.GetOneRequest()
		if err != nil {
			log.Println(err.Error())
			continue
		}

		log.Printf("get one request :%v", m)

		b, _ := hex.DecodeString(m.Info)
		err = comm.Writer(b)
		if err != nil {
			log.Println(err.Error())
			continue
		}

		m.DeleteMessage()
		log.Println("the request is deleted.")

		time.Sleep(time.Millisecond * t)

	}

}

func SessionRequest() {
	var r ftprotocol.RequestSession
	r.DeviceID = 0xD8
	r.ProtocolVersion = 0x2729
	r.SessionKey = []byte{0x46, 0x46}
	r.Sequence = 0x30
	r.NoAck = true

	var m conn.Message
	m.Messagetype = conn.REQUEST
	m.Info = hex.EncodeToString(r.Message())
	m.Status = conn.NONE

	m.InsertMessage()
}

func worker(t time.Duration) {

	for {
		log.Println("working....")
		time.Sleep(time.Millisecond * t)
	}

}

func reader(t time.Duration) {
	for {
		var m conn.Message
		b, _ := comm.Reader()
		if len(b) > 0 {
			m.Messagetype = conn.RESPONSE
			m.Info = hex.EncodeToString(b)
			m.Status = conn.NONE

			m.InsertMessage()

		}
		time.Sleep(time.Millisecond * t)

	}

}

func done(t time.Duration, wg *sync.WaitGroup) {
	bStop := false
	for {
		if bStop {
			for i := 0; i < len(funclist); i++ {
				wg.Done()
			}
			log.Println("done........")

		}
		time.Sleep(time.Millisecond * t)
	}

}

var funclist []string = []string{"generator", "writer", "reader", "worker", "done"}

func Test_concurrency() {
	err := comm.OpenSerial()

	if err != nil {
		log.Fatal(err)
	}
	var wg sync.WaitGroup
	wg.Add(len(funclist))

	go generator(time.Duration(250))
	go writer(time.Duration(200))
	go reader(time.Duration(200))
	go worker(time.Duration(1000))
	go done(time.Duration(1000000), &wg)

	wg.Wait()

}

func main() {
	//	Test_channel()

	Test_concurrency()

}
