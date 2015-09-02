package ftprotocol

import (
	"bytes"
	"encoding/binary"
	_ "log"
)

type GetSensor struct {
	Frame
	IsBroadcast     bool
	Broadcastperiod uint32
	IsAllSensorData bool
}

func (r GetSensor) ToByte() []byte {
	buf := new(bytes.Buffer)

	if r.IsBroadcast {
		binary.Write(buf, binary.BigEndian, uint32(1))
	} else {
		binary.Write(buf, binary.BigEndian, uint32(0))
	}
	binary.Write(buf, binary.BigEndian, r.Broadcastperiod)

	if r.IsAllSensorData {
		binary.Write(buf, binary.BigEndian, uint32(1))
	} else {
		binary.Write(buf, binary.BigEndian, uint32(0))
	}
	return buf.Bytes()
}

func (r GetSensor) Message() []byte {
	r.Frame.Init()

	r.MessageID = []byte{0x33, 0x42}

	r.MessageData = r.ToByte()

	//	log.Printf("%X", r.ByteArray())

	return r.ByteArray()
}
