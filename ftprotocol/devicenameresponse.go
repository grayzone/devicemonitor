package ftprotocol

import (
	"encoding/binary"
	"errors"
	"log"
	"strings"
)

type DeviceNameResponse struct {
	Frame
	SizeOfStringName uint32
	StringName       string
	NullTermination  uint32
}

func (r *DeviceNameResponse) Parse(input []byte) {
	r.Frame.Parse(input)
	data := r.Frame.MessageData
	r.ParseMessageData(data)
}

func (r *DeviceNameResponse) ParseMessageData(data []byte) error {
	log.Println("DeviceNameResponse message data length : ", len(data))
	if len(data) != 72 {
		return errors.New("DeviceNameResponse.ParseMessageData:invalid length.")
	}
	//	buf := bytes.NewReader(data)
	r.SizeOfStringName = binary.BigEndian.Uint32(data[:4])
	var name []byte
	for i := 67; i >= 4; i-- {
		if data[i] != 0x00 {
			name = data[4 : i+1]
			break
		} else {
			// find the unused characters

		}
	}
	r.StringName = strings.TrimSpace(string(name))
	r.NullTermination = binary.BigEndian.Uint32(data[68:72])

	log.Println("string size :", r.SizeOfStringName)
	log.Printf("device name : |%s|\n", r.StringName)
	log.Println("termination : ", r.NullTermination)

	//	log.Println(r)

	return nil
}
