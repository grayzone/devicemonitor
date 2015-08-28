package util

import (
	"fmt"
	_ "log"
)

var crc16Table []uint16 = make([]uint16, 256)

var crc_init bool = false

func InitCrc16() {
	var crc uint16
	for i, _ := range crc16Table {
		crc = uint16(i << 8)
		for j := 0; j < 8; j++ {
			var tmp uint16
			if (crc & 0x8000) == 0 {
				tmp = 0
			} else {
				tmp = 0x1021
			}
			crc = (crc << 1) ^ tmp
			crc16Table[i] = crc & 0xFFFF
		}
	}
	crc_init = true
}

func Crc16(addr []byte) uint16 {
	if crc_init != true {
		InitCrc16()
	}
	var crc uint16

	for i := 0; i < len(addr); i++ {
		crc = uint16(crc16Table[((crc>>8)&255)]) ^ (crc << 8) ^ (uint16(addr[i]) & 0x00FF)
		crc = crc & 0x0000FFFF
	}
	return crc
}

func Crc16Byte(addr []byte) []byte {
	//	log.Printf("%X", addr)
	crc := Crc16(addr)
	//	log.Printf("%X : %04X", addr, crc)
	s := fmt.Sprintf("%04X", crc)
	return []byte(s)
}
