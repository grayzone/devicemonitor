package ftprotocol

import (
	//	"bytes"
	"encoding/hex"
	"errors"
	"reflect"
	"testing"
)

func TestByteArray(t *testing.T) {
	cases := []struct {
		in   Frame
		want string
	}{
		{Frame{NoAck: true}, "00003030303000"},
		{Frame{Start: STX, End: ETX}, "060002003032303003"},
	}

	for _, c := range cases {
		//		c.in.NoAck = true
		output := c.in.ByteArray()

		got := hex.EncodeToString(output)

		if c.want != got {
			t.Errorf("ByteArray(),\nwant\t%s\ngot\t%s\n", c.want, got)
		}
	}
}

func TestParse(t *testing.T) {
	cases := []struct {
		in    string
		want1 string
		want2 string
		want3 error
	}{
		{"0630", "", "", errors.New("no message parsed")},
		{"06", "", "", errors.New("invalid length")},
		{"06309", "", "", errors.New("invalid length")},
		{"0245353231323030202020202020202029524420202020202020203342202020202f4820202020262020202020202020202020314630370306300630", "06300630", "3132", nil},
		{"06300245353231323030202020202020202029524420202020202020203342202020202f48202020202620202020202020202020203146303703", "", "3132", nil},
		{"06300245353231323030202020202020202029524420202020202020203342202020202f482020202026202020202020202020202031463037030630", "0630", "3132", nil},
		{"0245353231323030202020202020202029524420202020202020203342202020202f4820202020262020202020202020202020314630370306300245353231323030202020202020202029524420202020202020203342202020202f482020202026202020202020202020202031463037030630", "06300245353231323030202020202020202029524420202020202020203342202020202f482020202026202020202020202020202031463037030630", "3132", nil},
		{"06300238353031323030202020202020202029524420202020202020203342202020202f482020202026202020202020202020202046373036030630", "0630", "3132", nil},
	}
	for _, c := range cases {
		var got Frame
		input, _ := hex.DecodeString(c.in)
		got1, got3 := got.Parse(input)
		got2 := hex.EncodeToString(got.MessageID)
		if got1 != c.want1 || got2 != c.want2 || !reflect.DeepEqual(got3, c.want3) {
			t.Errorf("Parse(),\nwant\t|%s|\t%s\t%v\ngot\t|%s|\t%s\t%v\n", c.want1, c.want2, c.want3, got1, got2, got3)

		}

	}

}
