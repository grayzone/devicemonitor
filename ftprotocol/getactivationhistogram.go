package ftprotocol

type GetActivationHistogram struct {
	Frame
}

func (g GetActivationHistogram) Message() []byte {
	g.Frame.Init()

	g.MessageData = nil

	return g.ByteArray()
}
