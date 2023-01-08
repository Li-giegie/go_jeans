package go_jeans

import (
	"encoding/json"
)

func (a *MessageA) Reply(msg []byte)([]byte,error)  {
	a.Msg = msg
	return _Pack(a)
}

func (a *MessageA) Reply_String(msg string)  ([]byte,error) {
	return a.Reply([]byte(msg))
}

func (a *MessageA) Reply_JSON(obj interface{}) ([]byte,error) {
	buf,err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	return a.Reply(buf)
}

func (b *MessageB) Reply(msg []byte) ([]byte,error) {
	b.Msg = msg
	return _Pack(b)
}

func (b *MessageB) Reply_String(msg string) ([]byte,error) {
	return b.Reply([]byte(msg))
}

func (b *MessageB) Reply_JSON(obj interface{}) ([]byte,error)  {
	buf,err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	return b.Reply(buf)
}