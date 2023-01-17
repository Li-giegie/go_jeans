package go_jeans

type MessageI interface {
	Marshal () ([]byte,error)
	Unmarshal () (MessageI,error)
	Reply (msg []byte)
	GetMessage () interface{}
}

