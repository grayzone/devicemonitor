package ftprotocol

type KeepAlive struct {
	Frame
}

func (ka KeepAlive) ToByte() []byte {
	return nil
}

func (ka KeepAlive) Message() []byte {
	ka.Start = STX
	ka.SessionKey = []byte("FF")
	ka.Sequence = byte('2')
	ka.MessageID = []byte("00")
	ka.Unused = []byte("00")

	ka.MessageData = nil
	ka.End = ETX

	return ka.ByteArray()
}
