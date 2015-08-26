package ftprotocol

type GetRunTime struct {
	Frame
}

func (g GetRunTime) Message() []byte {
	g.Frame.Init()

	g.MessageData = nil
	//  ka.End = ETX

	return g.ByteArray()
}
