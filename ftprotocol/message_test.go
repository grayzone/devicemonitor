package ftprotocol

import (
	"encoding/hex"
	"errors"
	"reflect"
	"testing"
)

func TestMessage(t *testing.T) {
	cases := []struct {
		in   Message
		want string
	}{
		{RequestSession{Frame: Frame{SessionKey: []byte("FF"), Sequence: byte('0'), MessageID: []byte("11")}, DeviceID: DeviceID, ProtocolVersion: ProtocolVer}, "024646303131303020202020203020202952402020202020202020202020202020202020202020202020202020203536313103"},
		{RequestSession{Frame: Frame{SessionKey: []byte("FF"), Sequence: byte('2'), MessageID: []byte("11")}, DeviceID: 0x01, ProtocolVersion: 0x2727}, "0246463231313030202020202030202029523c2020202020202020202020202020202020202020202020202020204231334303"},
		{KeepAlive{Frame{SessionKey: []byte("FF"), Sequence: byte('0'), MessageID: []byte("00")}}, "02464630303030304438464203"},
	}

	for _, c := range cases {

		output := c.in.Message()
		got := hex.EncodeToString(output)

		if got != c.want {
			t.Errorf("Message(),\nwant\t%s\ngot\t%s\n", c.want, got)
		}
	}
}

func TestGetMessageID(t *testing.T) {
	cases := []struct {
		in    string
		want1 int
		want2 error
	}{
		{"024646303131303020202020203020202952402020202020202020202020202020202020202020202020202020203536313103", 0x11, nil},
		{"024646353132303020202020202020202952402020202020202020235a202020202f48202020202320202020202020202020203634413103", 0x12, nil},
		{"0246463", 0x00, errors.New("invalid message length")},
	}

	for _, c := range cases {
		input, _ := hex.DecodeString(c.in)
		got1, got2 := GetMessageID(input)

		if got1 != c.want1 || !reflect.DeepEqual(got2, c.want2) {
			t.Errorf("Message(),\ngiven\t%s\nwant\t%X:%v\ngot\t%X:%v\n", c.in, c.want1, c.want2, got1, got2)
		}
	}
}

func TestToByte(t *testing.T) {
	cases := []struct {
		in   Message
		want string
	}{
		//	{RequestSession{DeviceID: DeviceID, ProtocolVersion: ProtocolVer}, "00000001000027280000000000000000000000000000000000000000"},
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
		{0x11, MessageTable{"RequestSession", 0x11, Encoded, 56}},
		{0x00, MessageTable{"KeepAlive", 0x00, ASCII, 0}},
	}

	for _, c := range cases {
		got := FindMessageTable(c.in)
		if !reflect.DeepEqual(*got, c.want) {
			//if !bytes.Equal(got, c.want) {
			t.Errorf("FindMessageTable()\nwant\t%v\ngot\t%v", c.want, *got)
		}
	}
}

func TestParse(t *testing.T) {
	cases := []struct {
		in   string
		want MessageTable
	}{
		{"024646313132303020202020202020202952402020202020202020235a202020202f48202020202320202020202020202020204537303703", MessageTable{}},
	}

	for _, c := range cases {

		input, _ := hex.DecodeString(c.in)

		got := Parse(input)

		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("ByteArray(),\nwant\t%q\ngot\t%q\n", c.want, got)
		}
	}
}
