package ftprotocol

import (
	"bytes"
	"encoding/binary"
)

type ScreenPress struct {
	Frame
	ScreenID uint32
	X        uint32
	Y        uint32
}

func (r ScreenPress) ToByte() []byte {
	buf := new(bytes.Buffer)

	binary.Write(buf, binary.BigEndian, r.ScreenID)
	binary.Write(buf, binary.BigEndian, r.X)
	binary.Write(buf, binary.BigEndian, r.Y)

	return buf.Bytes()
}

func (r ScreenPress) Message() []byte {
	r.Frame.Init()

	r.MessageID = []byte{0x36, 0x32}

	r.MessageData = r.ToByte()

	//	log.Printf("%X", r.ByteArray())

	return r.ByteArray()
}
