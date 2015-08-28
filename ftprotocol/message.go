package ftprotocol

import (
	_ "encoding/hex"
	"errors"
	"github.com/grayzone/devicemonitor/util"
	"log"
	_ "reflect"
	"strconv"
)

type MessageEncoding int

const (
	Encoded = iota
	ASCII
)

const (
	REQUESTSESSION         = 0x11
	KEEPALIVE              = 0x00
	REQUESTSESSIONRESPONSE = 0x12
	GETRUNTIME             = 0x2D
	DEVICENAMEREQUEST      = 0x1D
	GETVERSIONSREQUEST     = 0x5A
)

type Message interface {
	ToByte() []byte
	Message() []byte
}

type MessageTable struct {
	Name     string
	ID       int
	Encoding MessageEncoding
	Length   int
}

var MessageList []MessageTable = []MessageTable{
	MessageTable{"RequestSession", 0x11, Encoded, 56},
	MessageTable{"KeepAlive", 0x00, ASCII, 0},
	MessageTable{"DeviceNameRequest", 0x1D, Encoded, 0},
	MessageTable{"RequestSessionResponse", 0x12, Encoded, 64},
	MessageTable{"GetSensor", 0x3B, Encoded, 24},
	MessageTable{"GetCriticalData", 0x37, Encoded, 64},
	MessageTable{"GetActivationHistogram", 0x2B, Encoded, 0},
}

func FindMessageTable(id int) *MessageTable {
	for _, v := range MessageList {
		if v.ID == id {
			return &v
		}
	}
	return nil
}

func GetMessageID(input []byte) (int, error) {
	if len(input) < 13 {
		return 0, errors.New("invalid message length")
	}
	messageid := string(input[4:6])
	//	log.Printf("%s", messageid)
	msgid, err := strconv.ParseInt(messageid, 16, 32)
	if err != nil {
		return 0, err
	}
	return int(msgid), nil
}

func Parse(input []byte) MessageTable {
	log.Println(len(input))
	log.Printf("%X", input)
	var result MessageTable
	if len(input) == 0 {
		return result
	}
	if input[0] == ACK {
		input = input[2:]
	}

	messageid := string(input[4:6])
	msgid, _ := strconv.ParseInt(messageid, 16, 32)
	msg := FindMessageTable(int(msgid))

	var msglen int
	if msg.Encoding == Encoded {
		msglen = util.UnEncodeLength(msg.Length)
	} else {
		msglen = msg.Length
	}

	log.Println(len(input)-13, msglen/2)

	log.Println("Message", msg, msgid)

	return result

}
