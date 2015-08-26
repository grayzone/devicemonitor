package ftprotocol

type DeviceNameRequest struct {
	Frame
}

func (d DeviceNameRequest) Message() []byte {
	d.Frame.Init()

	d.MessageData = nil
	return d.ByteArray()
}
