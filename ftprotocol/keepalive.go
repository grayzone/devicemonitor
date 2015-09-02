package ftprotocol

type KeepAlive struct {
	Frame
}

func (ka KeepAlive) ToByte() []byte {
	return nil
}

func (ka KeepAlive) Message() []byte {
	ka.Frame.Init()

	ka.MessageID = []byte{0x30, 0x30}

	ka.MessageData = nil
	//	ka.End = ETX

	return ka.ByteArray()
}
