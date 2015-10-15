package conn

import (
	"log"
	"reflect"
	"testing"
)

func TestUpdateStatus(t *testing.T) {
	cases := []struct {
		in   MessageStatus
		want MessageStatus
	}{
		{NONE, NONE},
		{PROCESSED, PROCESSED},
		{INVALID, INVALID},
	}
	for _, c := range cases {
		var m Message
		m.Status = c.in
		m.InsertMessage()
		m.Status = c.want
		m.UpdateStatus()
		m.Get()
		got := m.Status
		if m.Status != c.want {
			t.Errorf("UpdateStatus(),\nwant\t%v\ngot\t%v\n", c.want, got)
		}

		m.DeleteMessage()
	}

}

func TestGetSetting(t *testing.T) {

	var s Setting
	err := s.GetSetting()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("TestGetSetting")
	log.Printf("%v\n", s)
	log.Println(s.Isconnected)
	log.Println(s.Deviceid)
	log.Println(s.Protocolver)
	log.Println("...")
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

func TestGetSensordata(t *testing.T) {
	var s Sensordata
	s.GetSensordata()
	log.Println("TestGetSensordata")
	log.Println(s.Vavg)
	log.Println(s.Iavg)
	log.Println(s.Icf)
	log.Println(s.Createtime.String())
	log.Println("...")
}
