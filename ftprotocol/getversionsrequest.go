package ftprotocol

type GetVersionsRequest struct {
	Frame
}

func (g GetVersionsRequest) Message() []byte {
	g.Frame.Init()

	g.MessageData = nil
	return g.ByteArray()
}
