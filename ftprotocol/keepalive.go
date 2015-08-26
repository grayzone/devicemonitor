package ftprotocol

type KeepAlive struct {
	Frame
}

func (ka KeepAlive) ToByte() []byte {
	return nil
}

func (ka KeepAlive) Message() []byte {
	ka.Frame.Init()

	ka.MessageData = nil
	//	ka.End = ETX

	return ka.ByteArray()
}
