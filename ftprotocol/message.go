package ftprotocol

type MessageEncoding int

const (
	Encoded = iota
	ASCII
)

const (
	REQUESTSESSION = 0x11
	KEEPALIVE      = 0x00
)

type Message interface {
	ToByte() []byte
	Message() []byte
}

type MessageTable struct {
	Name     string
	ID       int
	Encoding MessageEncoding
	Type     interface{}
}

var MessageList []MessageTable = []MessageTable{
	MessageTable{"RequestSession", 0x11, Encoded, RequestSession{}},
	MessageTable{"KeepAlive", 0x00, ASCII, KeepAlive{}},
}

func FindMessageTable(id int) *MessageTable {
	for _, v := range MessageList {
		if v.ID == id {
			return &v
		}
	}
	return nil
}
