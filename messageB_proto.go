package go_jeans

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"google.golang.org/protobuf/proto"
	"io"
)

func NewMsgB_Proto(msg []byte,SrcAddr,DestApi,DestAddr uint32) *MessageB_Proto  {
	lock.Lock()
	count++
	defer lock.Unlock()
	return &MessageB_Proto{
		MsgId:         count,
		Msg:           msg,
		SrcAddr:       SrcAddr,
		DestApi:       DestApi,
		DestAddr:      DestAddr,
	}
}

func (a *MessageB_Proto) Marshal() ([]byte,error) {
	var buf = new(bytes.Buffer)
	pbuf,err := proto.Marshal(a)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(),binary.Write(buf,binary.LittleEndian,uint32(len(pbuf)))
}

func (a *MessageB_Proto) Unmarshal(conn io.Reader) (*MessageB,error) {
	buf,err := _read(conn)
	if err != nil {
		return nil,err
	}
	var tmp = new(MessageB)

	return tmp,json.Unmarshal(buf,&tmp)
}

func (a *MessageB_Proto) Reply(msg []byte) ([]byte,error) {
	a.Msg = msg
	return a.Marshal()
}