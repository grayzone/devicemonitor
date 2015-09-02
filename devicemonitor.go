package main

import (
	"encoding/hex"
	"log"
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

func generatoer() {
	var sequence byte = 0x30
	for {
		var k ftprotocol.KeepAlive
		k.SessionKey = []byte{0x46, 0x46}
		k.Sequence = sequence
		log.Printf("send keepalive : %X", k.Message())

		var m conn.Message
		m.Messagetype = conn.REQUEST
		m.Info = hex.EncodeToString(k.Message())
		m.Status = conn.NONE

		m.InsertMessage()

		sequence = IncreaseSeq(sequence)

		time.Sleep(time.Millisecond * 200)

	}

}

func writer() {
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
		comm.Writer(b)

		m.DeleteMessage()
		log.Println("the request is deleted.")

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

func worker() {

}

func reader() {
	for {
		var m conn.Message
		b, _ := comm.Reader()
		if len(b) > 0 {
			m.Messagetype = conn.RESPONSE
			m.Info = hex.EncodeToString(b)
			m.Status = conn.NONE

			m.InsertMessage()
		}

	}

}

func Test_channel() {

	err := comm.OpenSerial()
	if err != nil {
		log.Println(err.Error())
	}
	go generatoer()
	go writer()
	go reader()
	go worker()
	go SessionRequest()
	time.Sleep(time.Second * 10)
}

func main() {
	Test_channel()

}
