package comm

import (
	"errors"
	"log"
	"time"

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
	c.ReadTimeout = time.Millisecond * 200
	var err error
	s = new(serial.Port)
	s, err = serial.OpenPort(c)
	return err
}

func CloseSerial() error {
	if s != nil {
		return s.Close()
	}
	log.Println("serial port is closed.")
	return nil
}

//c.Baud = 115200
//c.Name = "COM1"
//c.ReadTimeout = time.Millisecond*1000
//c := &serial.Config{Name: "COM1", Baud: 115200, ReadTimeout: time.Millisecond * 1000}

func Sender(msg []byte) []byte {

	//	s, _ := serial.OpenPort(c)

	log.Printf("SEND: %X", msg)

	//	defer s.Close()
	var n int
	n, err := s.Write(msg)
	if err != nil {
		log.Fatal(err)
	}
	//	time.Sleep(time.Millisecond * 100)
	result := make([]byte, 2048)
	n, err = s.Read(result)
	if err != nil {
		log.Fatal(err)
	}
	return result[:n]

}

func Writer(msg []byte) error {
	n, err := s.Write(msg)
	log.Println(n, err)
	if n == 0 {
		return errors.New("failed to write data to device.")
	}
	return err
}

func Reader() ([]byte, error) {
	result := make([]byte, 2048)

	n, err := s.Read(result)
	if err != nil {
		return []byte(""), err
	}
	return result[:n], nil
}
