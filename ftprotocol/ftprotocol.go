package ftprotocol

import (
	"bytes"
	"strconv"

	//	"encoding/hex"
	"github.com/grayzone/devicemonitor/util"
	_ "log"
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

	f.CRC = util.Crc16Byte(result.Bytes())

	result.Write(f.CRC)

	result.WriteByte(f.End)

	return result.Bytes()
}

func (f *Frame) Parse([]byte) {

}
