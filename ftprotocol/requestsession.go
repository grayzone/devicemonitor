package ftprotocol

import (
	"bytes"
	"encoding/binary"
	_ "log"
)

type RequestSession struct {
	Frame
	DeviceID        uint32
	ProtocolVersion uint32
	Reserved1       uint32
	Reserved2       uint32
	Reserved3       uint32
	Reserved4       uint32
	Reserved5       uint32
}

func (r RequestSession) ToByte() []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, r.DeviceID)
	binary.Write(buf, binary.BigEndian, r.ProtocolVersion)
	binary.Write(buf, binary.BigEndian, r.Reserved1)
	binary.Write(buf, binary.BigEndian, r.Reserved2)
	binary.Write(buf, binary.BigEndian, r.Reserved3)
	binary.Write(buf, binary.BigEndian, r.Reserved4)
	binary.Write(buf, binary.BigEndian, r.Reserved5)

	return buf.Bytes()
}

func (r RequestSession) Message() []byte {
	r.Frame.Init()
	r.MessageID = []byte{0x31, 0x31}

	//	r.DeviceID = DeviceID
	//	r.ProtocolVersion = ProtocolVer
	r.MessageData = r.ToByte()

	//	log.Printf("%X", r.ByteArray())

	return r.ByteArray()
}
