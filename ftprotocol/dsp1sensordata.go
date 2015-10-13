package ftprotocol

import (
	"bytes"
	"encoding/binary"
	"errors"
	"log"
)

type Dsp1SensorData struct {
	Frame
	IsValidData      bool
	SequenceNumber   uint32
	IsActivatingFlag bool
	Vavg             float32
	Iavg             float32
	Pavg             float32
	Vrms             float32
	Irms             float32
	Viphase          float32
	Vpk              float32
	Ipk              float32
	Vcf              float32
	Icf              float32
	Zload            float32
	T1               float32
	T2               float32
	Leakage          float32
	Stimpos          float32
	Stimneg          float32
	Oltarget         float32
	//	Stimneg2         float32
}

func (r *Dsp1SensorData) Parse(input []byte) {
	r.Frame.Parse(input)
	data := r.Frame.MessageData
	r.ParseMessageData(data)
}

func readbytetofloat(b []byte) float32 {
	var result float32
	buf := bytes.NewReader(b)
	err := binary.Read(buf, binary.BigEndian, &result)
	if err != nil {
		log.Printf("%X, binary read failed :%s", b, err.Error())
		result = 0.0
	}
	return result
}

func (r *Dsp1SensorData) ParseMessageData(data []byte) error {
	log.Println("Dsp1SensorData message data length : ", len(data))
	if len(data) != 80 {
		return errors.New("Dsp1SensorData.ParseMessageData:invalid length.")
	}

	var iisvalidata uint32
	iisvalidata = binary.BigEndian.Uint32(data[:4])
	if iisvalidata == 0 {
		r.IsValidData = false
	} else {
		r.IsValidData = true
	}

	r.SequenceNumber = binary.BigEndian.Uint32(data[4:8])

	var iisactivatingflag uint32
	iisactivatingflag = binary.BigEndian.Uint32(data[8:12])
	if iisactivatingflag == 0 {
		r.IsActivatingFlag = false
	} else {
		r.IsActivatingFlag = true
	}

	r.Vavg = readbytetofloat(data[12:16])
	r.Iavg = readbytetofloat(data[16:20])
	r.Pavg = readbytetofloat(data[20:24])
	r.Vrms = readbytetofloat(data[24:28])
	r.Irms = readbytetofloat(data[28:32])
	r.Viphase = readbytetofloat(data[32:36])
	r.Vpk = readbytetofloat(data[36:40])
	r.Ipk = readbytetofloat(data[40:44])
	r.Vcf = readbytetofloat(data[44:48])
	r.Icf = readbytetofloat(data[48:52])
	r.Zload = readbytetofloat(data[52:56])
	r.T1 = readbytetofloat(data[56:60])
	r.T2 = readbytetofloat(data[60:64])
	r.Leakage = readbytetofloat(data[64:68])
	r.Stimpos = readbytetofloat(data[68:72])
	r.Stimneg = readbytetofloat(data[72:76])
	r.Oltarget = readbytetofloat(data[76:80])
	//	r.Stimneg2 = readbytetofloat(data[80:84])

	log.Println("IsValidData : ", r.IsValidData)
	log.Println("Sequence number : ", r.SequenceNumber)
	log.Println("IsActivatingFlag : ", r.IsActivatingFlag)

	log.Println("Vavg : ", r.Vavg)
	log.Println("Iavg : ", r.Iavg)
	log.Println("Pavg : ", r.Pavg)
	log.Println("Vrms : ", r.Vrms)
	log.Println("Irms : ", r.Irms)
	log.Println("Viphase : ", r.Viphase)
	log.Println("Vpk : ", r.Vpk)
	log.Println("Ipk : ", r.Ipk)
	log.Println("Vcf : ", r.Vcf)
	log.Println("Icf : ", r.Icf)
	log.Println("Zload : ", r.Zload)
	log.Println("T1 : ", r.T1)
	log.Println("T2 : ", r.T2)
	log.Println("Leakage : ", r.Leakage)
	log.Println("Stimpos : ", r.Stimpos)
	log.Println("Stimneg : ", r.Stimneg)
	log.Println("Oltarget : ", r.Oltarget)
	//	log.Println("Stimneg2 : ", r.Stimneg2)

	return nil
}
