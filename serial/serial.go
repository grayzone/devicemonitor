package serial

import (
	"log"
	"time"

	"github.com/tarm/serial"
)

func Sender(msg []byte) []byte {
	log.Printf("%X", msg)
	c := &serial.Config{Name: "COM1", Baud: 115200, ReadTimeout: time.Second * 5}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}
	defer s.Close()
	var n int
	n, err = s.Write(msg)
	if err != nil {
		log.Fatal(err)
	}
	result := make([]byte, 10240)
	n, err = s.Read(result)
	if err != nil {
		log.Fatal(err)
	}
	return result[:n]

}
