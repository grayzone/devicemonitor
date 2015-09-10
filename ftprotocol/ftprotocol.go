package ftprotocol

import (
	"bytes"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/grayzone/devicemonitor/util"
)

var ProtocolVer uint32 = 0x2728
var DeviceID uint32 = 0x01

const (
	ACK   byte = 0x06
	CR    byte = 0x0D
	ETX   byte = 0x03
	LF    byte = 0x0A
	NAK   byte = 0x15
	SPACE byte = 0x20
	STX   byte = 0x02
)

type Frame struct {
	Start       byte
	SessionKey  []byte
	Sequence    byte
	MessageID   []byte
	Unused      []byte
	MessageData []byte
	CRC         []byte
	End         byte
	NoAck       bool
}

func (f *Frame) Init() {
	f.Start = STX
	f.Unused = []byte("00")
	f.End = ETX
}

func (f *Frame) ByteArray() []byte {

	sid := string(f.MessageID)
	//	log.Println(sid)
	msgid, _ := strconv.ParseInt(sid, 16, 64)

	//	log.Println(msgid)
	msg := FindMessageTable(int(msgid))

	result := new(bytes.Buffer)
	/*
		if !bytes.Equal(f.MessageID, []byte{0x31, 0x31}) {
			result.WriteByte(ACK)
			result.WriteByte(f.Sequence)
		}
	*/
	if !f.NoAck {
		result.WriteByte(ACK)
		result.WriteByte(f.Sequence)

	}

	result.WriteByte(f.Start)
	result.Write(f.SessionKey)
	result.WriteByte(f.Sequence)
	result.Write(f.MessageID)

	result.Write(f.Unused)

	if f.MessageData != nil {
		if msg.Encoding == Encoded {
			result.Write(util.UuEnocde(f.MessageData))
		} else {
			result.Write(f.MessageData)
		}

	}
	if !f.NoAck {
		f.CRC = util.Crc16Byte(result.Bytes()[2:])
	} else {
		f.CRC = util.Crc16Byte(result.Bytes())
	}

	result.Write(f.CRC)

	result.WriteByte(f.End)

	return result.Bytes()
}

func (f *Frame) Parse(input []byte) (string, error) {
	var result string
	if len(input) < 2 {
		return result, errors.New("invalid length")
	}
	for {
		if len(input) >= 2 {
			if input[0] == ACK || input[0] == NAK {
				if input[1] <= 0x39 && input[1] >= 0x30 {
					input = input[2:]
				}
			} else {
				break
			}
		} else {
			return result, errors.New("no message parsed")
		}
	}
	if len(input) < 13 {
		return hex.EncodeToString(input), errors.New("incomplete input")
	}
	// get message id
	messageid := string(input[4:6])
	msgid, _ := strconv.ParseInt(messageid, 16, 32)
	msg := FindMessageTable(int(msgid))
	if msg == nil {
		errmsg := fmt.Sprintf("not found the mssage id : %d", msgid)
		return "", errors.New(errmsg)
	}

	var msglen int
	if msg.Encoding == Encoded {
		msglen = util.UnEncodeLength(msg.Length)
	} else {
		msglen = msg.Length
	}
	msglen = msglen / 2
	//	log.Println("message length:", msglen)

	f.Start = input[0]
	f.SessionKey = input[1:3]
	f.Sequence = input[3]
	f.MessageID = input[4:6]
	f.Unused = input[6:8]
	if msg.Encoding != Encoded {
		f.MessageData = input[8 : msglen+8]
	} else {
		f.MessageData = util.UuDecode(input[8 : msglen+8])
	}

	//	f.MessageData = util.UuDecode(f.MessageData)
	f.CRC = input[msglen+8 : msglen+12]
	f.End = input[msglen+12]

	result = hex.EncodeToString(input[msglen+13:])

	log.Printf("start : %x\n", f.Start)
	log.Printf("SessionKey : %x\n", f.SessionKey)
	log.Printf("Sequence : %x\n", f.Sequence)
	log.Printf("MessageID : %x\n", f.MessageID)
	log.Printf("Unused : %x\n", f.Unused)
	log.Printf("MessageData : %x\n", f.MessageData)
	log.Printf("CRC : %x\n", f.CRC)
	log.Printf("End : %x\n", f.End)

	return result, nil

}
