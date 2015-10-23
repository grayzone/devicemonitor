package main

import (
	"encoding/hex"
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/grayzone/devicemonitor/comm"
	"github.com/grayzone/devicemonitor/conn"
	"github.com/grayzone/devicemonitor/ftprotocol"
)

var bSoftDelete = false

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
	for {
		var s conn.Setting
		s.GetSetting()
		if s.Isconnected {
			var k ftprotocol.KeepAlive
			k.SessionKey = []byte(s.Sessionkey)
			k.Sequence = byte(s.Sequence[0])

			//			log.Printf("session key : %X, sequence : %X\n", k.SessionKey, k.Sequence)

			//		log.Printf("send keepalive : %X", k.Message())

			var m conn.Message
			m.Messagetype = conn.REQUEST
			m.Info = hex.EncodeToString(k.Message())
			m.Status = conn.NONE

			m.InsertMessage()

			//			sequence = IncreaseSeq(sequence)

		}

		time.Sleep(time.Millisecond * time.Duration(t))
		//		time.Sleep(time.Millisecond * time.Duration(s.Writeinterval))

	}

}

func IncreaseOneSequence() {
	var s conn.Setting
	s.GetSetting()
	//	log.Printf(" seq  :%x", byte(s.Sequence[0]))
	s.Sequence = string(IncreaseSeq(byte(s.Sequence[0])))
	//	log.Printf("Sequence : %s", s.Sequence)
	s.UpdateSequence()
}

func writer(t time.Duration) {
	// get one request
	for {
		var s conn.Setting
		s.GetSetting()
		//		time.Sleep(time.Millisecond * time.Duration(s.Writeinterval))
		time.Sleep(time.Millisecond * time.Duration(t))
		if s.Isconnected {
			var m conn.Message
			err := m.GetOneRequest()
			if err != nil {
				//				log.Println("writer:", err.Error())
				continue
			}

			//			log.Printf("get one request :%v", m)

			b, _ := hex.DecodeString(m.Info)
			err = comm.Writer(b)
			if err != nil {
				log.Println(err.Error())
				continue
			}
			log.Printf("Send:%X", b)
			IncreaseOneSequence()

			if bSoftDelete {
				m.Status = conn.DELETED
				m.UpdateStatus()
			} else {
				m.DeleteMessage()
				//		log.Println("the request is deleted.")
			}
		}

	}
}

func worker(t time.Duration) {

	for {
		time.Sleep(time.Millisecond * t)

		//		log.Println("working....")
		var m conn.Message
		err := m.GetOneResponse()
		if err != nil {
			//		log.Println("worker:", err.Error())
			continue
			//			log.Printf("get one response :%v", m)
		}

		//		log.Println(m.Info)

		b, err := hex.DecodeString(m.Info)
		if err != nil {
			log.Println(err.Error())
			m.Status = conn.INVALID
			m.UpdateStatus()
			continue
		}
		var f ftprotocol.Frame
		//		log.Printf("%X", b)
		s, err := f.Parse(b)
		if err != nil {
			log.Printf("%X:%s\n", b, err.Error())
			//	m.DeleteMessage()
			m.Status = conn.INVALID
			m.UpdateStatus()

			continue
		}

		msgid, _ := strconv.ParseInt(string(f.MessageID), 16, 32)

		switch msgid {
		case ftprotocol.REQUESTSESSIONRESPONSE:
			var sessionres ftprotocol.RequestSessionResponse
			sessionres.Frame = f
			err := sessionres.ParseMessageData(f.MessageData)
			if err != nil {
				log.Println(err.Error())
				m.Status = conn.INVALID
				m.UpdateStatus()
				continue
			}

			var s conn.Setting
			s.Sessionstatus = sessionres.SessionStatus
			s.UpdateSessionStatus()

			s.Sessiontimeout = sessionres.SessionTimeout
			s.UpdateSessiontimeout()

			s.Deviceid = conn.DeviceID(sessionres.DeviceID)
			s.UpdateDeviceid()

			s.Protocolver = strconv.FormatUint(uint64(sessionres.ProtocolVersion), 10)
			s.UpdateProtocolVer()

			s.Sessionkey = string(sessionres.SessionKey)
			s.UpdateSessionKey()

			s.Sequence = string(sessionres.Sequence)
			s.UpdateSequence()

			s.Messagetimeout = sessionres.MessageTimeout
			s.UpdateMessagetimeout()

			s.Maxretrycount = sessionres.MaxRetryCount
			s.UpdateMaxretrycount()

		case ftprotocol.DEVICENAMERESPONSE:
			var res ftprotocol.DeviceNameResponse
			res.Frame = f
			err := res.ParseMessageData(f.MessageData)
			if err != nil {
				log.Println(err.Error())
				m.Status = conn.INVALID
				m.UpdateStatus()
				continue
			}
			var s conn.Setting
			//		log.Println("device  name : ", res.StringName)
			s.Devicename = res.StringName
			err = s.UpdateDevicename()
			if err != nil {
				log.Println("UpdateDevicename:", err.Error())
			}

		case ftprotocol.GETSENSORRESPONSE:
			var res ftprotocol.GetSensorResponse
			res.Frame = f
			err := res.ParseMessageData(f.MessageData)
			if err != nil {
				log.Println("GETSENSORRESPONSE:", err.Error())
				m.Status = conn.INVALID
				m.UpdateStatus()
				continue
			}
			var s conn.Setting
			s.Sensorbroadcastperiod = res.BroadcastPeriod
			err = s.UpdateSensorbroadcastperiod()
			if err != nil {
				log.Println("UpdateSensorbroadcastperiod :", err.Error())
			}

		case ftprotocol.DSP1SENSORDATA:
			var res ftprotocol.Dsp1SensorData
			res.Frame = f
			err := res.ParseMessageData(f.MessageData)
			if err != nil {
				log.Println("DSP1SENSORDATA:", err.Error())
				m.Status = conn.INVALID
				m.UpdateStatus()
				continue
			}
			var s conn.Sensordata
			s.Isvaliddata = res.IsValidData
			s.Sequencenumber = res.SequenceNumber
			s.Isactivatingflag = res.IsActivatingFlag
			s.Vavg = res.Vavg
			s.Iavg = res.Iavg
			s.Pavg = res.Pavg
			s.Vrms = res.Vrms
			s.Irms = res.Irms
			s.Viphase = res.Viphase
			s.Vpk = res.Vpk
			s.Ipk = res.Ipk
			s.Vcf = res.Vcf
			s.Icf = res.Icf
			s.Zload = res.Zload
			s.T1 = res.T1
			s.T2 = res.T2
			s.Leakage = res.Leakage
			s.Stimpos = res.Stimpos
			s.Stimneg = res.Stimneg
			s.Oltarget = res.Oltarget

			s.Createtime = time.Now()

			s.InsertSensordata()

		case ftprotocol.EMPTY:
			var s conn.Setting
			s.Sequence = string(f.Sequence)
			//			log.Printf("sequence : %X", s.Sequence)
			s.UpdateSequence()
		default:
			log.Printf("unsupported message : %X", msgid)

		}
		if len(s) > 0 {
			m.Info = s
			m.UpdateInfo()
		} else {
			//			log.Println(m)
			if bSoftDelete {
				m.Status = conn.DELETED
				m.UpdateStatus()
			} else {
				m.DeleteMessage()
				//		log.Println("the request is deleted.")
			}
		}

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
			log.Printf("Received : %X", b)
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
		log.Fatal("Test_concurrency:", err)
	}
	var wg sync.WaitGroup
	wg.Add(len(funclist))

	go generator(time.Duration(200))
	go writer(time.Duration(1))
	go reader(time.Duration(1))
	go worker(time.Duration(1))
	//	go done(time.Duration(1000000), &wg)

	wg.Wait()

}

func main() {
	//	Test_channel()

	Test_concurrency()

}
