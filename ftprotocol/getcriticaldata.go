package ftprotocol

import (
	"bytes"
	"encoding/binary"
)

type GetCriticalData struct {
	Frame
	DataStoreNameSize uint32
	DataStoreName     string
}

func (r GetCriticalData) ToByte() []byte {
	buf := new(bytes.Buffer)

	binary.Write(buf, binary.BigEndian, r.DataStoreNameSize)
	//    binary.Write(w, order, data)
	//	binary.Write([]byte(r.DataStoreName))

	return buf.Bytes()
}

func (r GetCriticalData) Message() []byte {
	r.Frame.Init()

	r.MessageData = r.ToByte()

	
    //	log.Printf("%X", r.ByteArray())

	return r.ByteArray()
}
