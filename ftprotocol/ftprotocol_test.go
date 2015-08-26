package ftprotocol

import (
	//	"bytes"
	"encoding/hex"
	_ "reflect"
	"testing"
)

func TestByteArray(t *testing.T) {
	cases := []struct {
		in   Frame
		want string
	}{
		{Frame{}, "00003000"},
		{Frame{Start: STX, End: ETX}, "020032303003"},
	}

	for _, c := range cases {
		output := c.in.ByteArray()
		got := hex.EncodeToString(output)

		if c.want != got {
			t.Errorf("ByteArray(),\nwant\t%s\ngot\t%s\n", c.want, got)
		}
	}
}
