package connection

import (
	//	"bytes"
	"encoding/hex"
	"testing"
)

func TestSender(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		//		{"024646323131303020202020203020202952402020202020202020202020202020202020202020202020202020203337463003", ""},
		{"023030303131303020202020203020202952402020202020202020202020202020202020202020202020202020204131383303", ""},
		//{"02464632303030304236394203", ""},
	}

	for _, c := range cases {
		input, _ := hex.DecodeString(c.in)
		output := Sender(input)
		got := hex.EncodeToString(output)

		if got != c.want {
			t.Errorf("Sender()\nwant\t%s\ngot\t%s\n", c.want, got)
		}
	}
}
