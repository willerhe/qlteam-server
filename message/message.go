package message

type message interface {
	Payload() []byte
	Marshal(interface{})
	Unmarshal([]byte, interface{})
}

type Message struct {
	payload []byte
}
