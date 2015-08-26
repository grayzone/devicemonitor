package connection

import (
	"log"
	_ "time"

	"github.com/tarm/serial"
)

/*
func OpenPort() *Port {

}
*/
var c *serial.Config
var s *serial.Port

func OpenSerial() error {
	c = new(serial.Config)
	c.Baud = 115200
	c.Name = "COM1"
	//	c.ReadTimeout = time.Microsecond * 10
	var err error
	s = new(serial.Port)
	s, err = serial.OpenPort(c)
	return err
}

func CloseSerial() error {
	return s.Close()
}

//c.Baud = 115200
//c.Name = "COM1"
//c.ReadTimeout = time.Millisecond*1000
//c := &serial.Config{Name: "COM1", Baud: 115200, ReadTimeout: time.Millisecond * 1000}

func Sender(msg []byte) []byte {

	//	s, _ := serial.OpenPort(c)

	log.Printf("%X", msg)

	//	defer s.Close()
	var n int
	n, err := s.Write(msg)
	if err != nil {
		log.Fatal(err)
	}
	//	time.Sleep(time.Millisecond * 250)
	result := make([]byte, 2048)
	n, err = s.Read(result)
	if err != nil {
		log.Fatal(err)
	}
	return result[:n]

}
