package ftprotocol

import (
	"encoding/binary"
)

type RequestSessionResponse struct {
	Frame
	DeviceID        uint32
	ProtocolVersion uint32
	SessionStatus   uint32
	SessionTimeout  uint32
	MessageTimeout  uint32
	MaxRetryCount   uint32
	Reserved1       uint32
	Reserved2       uint32
}

func (r *RequestSessionResponse) Parse(input []byte) {
	r.Frame.Parse(input)
	data := r.Frame.MessageData
	r.DeviceID = binary.BigEndian.Uint32(data[:4])
	r.ProtocolVersion = binary.BigEndian.Uint32(data[4:8])
	r.SessionStatus = binary.BigEndian.Uint32(data[8:12])
	r.SessionTimeout = binary.BigEndian.Uint32(data[12:16])
	r.MessageTimeout = binary.BigEndian.Uint32(data[16:20])
	r.MaxRetryCount = binary.BigEndian.Uint32(data[20:24])
	r.Reserved1 = binary.BigEndian.Uint32(data[24:28])
	r.Reserved2 = binary.BigEndian.Uint32(data[28:32])

}
