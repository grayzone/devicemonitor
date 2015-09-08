package conn

import (
	"log"
	"reflect"
	"testing"
)

func TestGetSetting(t *testing.T) {

	var s Setting
	err := s.GetSetting()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%v\n", s)
	log.Println(s.Isconnected)
	log.Println(s.Deviceid)
	log.Println(s.Protocolver)
}

func TestUpdateIsConnected(t *testing.T) {
	cases := []struct {
		in   bool
		want bool
	}{
		{true, true},
		//		{false, false},
	}

	for _, c := range cases {
		var s Setting
		s.Isconnected = c.in
		err := s.UpdateIsConnected()
		if err != nil {
			t.Error(err)
		}
		var g Setting
		g.GetSetting()
		got := g.Isconnected

		if !reflect.DeepEqual(c.want, got) {
			t.Errorf("UpdateIsConnected(),\nwant\t%v\ngot\t%v\n", c.want, got)
		}
	}

}
