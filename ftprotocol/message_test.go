package ftprotocol

import (
	"encoding/hex"
	"reflect"
	"testing"
)

func TestMessage(t *testing.T) {
	cases := []struct {
		in   Message
		want string
	}{
		{RequestSession{Frame: Frame{SessionKey: []byte("FF"), Sequence: byte('0'), MessageID: []byte("11")}, DeviceID: 0x01, ProtocolVersion: 0x2727}, "024646303131303020202020203020202952402020202020202020202020202020202020202020202020202020203536313103"},
		{RequestSession{Frame: Frame{SessionKey: []byte("FF"), Sequence: byte('2'), MessageID: []byte("11")}, DeviceID: 0x01, ProtocolVersion: 0x2727}, "024646323131303020202020203020202952402020202020202020202020202020202020202020202020202020203337463003"},
		{KeepAlive{Frame{SessionKey: []byte("FF"), Sequence: byte('0'), MessageID: []byte("00")}}, "024646323131303020202020203020202952402020202020202020202020202020202020202020202020202020203337463003"},
	}

	for _, c := range cases {

		output := c.in.Message()
		got := hex.EncodeToString(output)

		if got != c.want {
			t.Errorf("Message(),\nwant\t%s\ngot\t%s\n", c.want, got)
		}
	}
}

func TestToByte(t *testing.T) {
	cases := []struct {
		in   Message
		want string
	}{
		{RequestSession{DeviceID: DeviceID, ProtocolVersion: ProtocolVer}, "00000001000027280000000000000000000000000000000000000000"},
		{RequestSession{DeviceID: 0x01, ProtocolVersion: 0x2727}, "00000001000027270000000000000000000000000000000000000000"},
		{KeepAlive{}, ""},
	}

	for _, c := range cases {
		output := c.in.ToByte()

		got := hex.EncodeToString(output)
		if got != c.want {
			//if !bytes.Equal(got, c.want) {
			t.Errorf("ToByte()\nwant\t%s\ngot\t%s", c.want, got)
		}
	}
}

func TestFindMessageTable(t *testing.T) {
	cases := []struct {
		in   int
		want MessageTable
	}{
		{0x11, MessageTable{"RequestSession", 0x11, Encoded, RequestSession{}}},
		{0x00, MessageTable{"KeepAlive", 0x00, ASCII, KeepAlive{}}},
	}

	for _, c := range cases {
		got := FindMessageTable(c.in)
		if !reflect.DeepEqual(*got, c.want) {
			//if !bytes.Equal(got, c.want) {
			t.Errorf("FindMessageTable()\nwant\t%v\ngot\t%v", c.want, *got)
		}
	}
}
