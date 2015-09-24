package ftprotocol

import (
	"encoding/binary"
	"errors"
	"log"
)

type GetSensorResponse struct {
	Frame
	BroadcastPeriod uint32
}

func (r *GetSensorResponse) Parse(input []byte) {
	r.Frame.Parse(input)
	data := r.Frame.MessageData
	r.ParseMessageData(data)
}

func (r *GetSensorResponse) ParseMessageData(data []byte) error {
	log.Println("GetSensorResponse message data length : ", len(data))
	if len(data) != 4 {
		return errors.New("GetSensorResponse.ParseMessageData:invalid length.")
	}

	r.BroadcastPeriod = binary.BigEndian.Uint32(data)

	log.Println("Broadcas tPeriod : ", r.BroadcastPeriod)

	//	log.Println(r)

	return nil
}
