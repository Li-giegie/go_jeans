package go_jeans

import (
	"bytes"
	"encoding/binary"
	"io"
)


type MessageB struct {
	MsgId         uint32 `json:"MsgId,omitempty"`
	Msg           []byte `json:"Msg,omitempty"`
	SrcAddr       uint32 `json:"SrcAddr,omitempty"`
	DestApi       uint32 `json:"DestApi,omitempty"`
	DestAddr      uint32 `json:"DestAddr,omitempty"`
}

func NewMsgB(msg []byte,SrcAddr,DestApi,DestAddr uint32) *MessageB  {
	lock.Lock()
	count++
	defer lock.Unlock()
	return &MessageB{
		MsgId:    count,
		Msg:      msg,
		SrcAddr:  0,
		DestApi:  0,
		DestAddr: 0,
	}
}

func (a *MessageB) Marshal() (*bytes.Buffer,error) {
	var buf = new(bytes.Buffer)

	err := binary.Write(buf,binary.LittleEndian,uint32(16+len(a.Msg)))
	if err != nil {
		return nil, err
	}
	if err = binary.Write(buf,binary.LittleEndian,a.MsgId); err != nil {
		return nil, err
	}
	if _,err = buf.Write(a.Msg);err != nil {
		return nil, err
	}
	if err = binary.Write(buf,binary.LittleEndian,a.SrcAddr); err != nil {
		return nil, err
	}

	if err = binary.Write(buf,binary.LittleEndian,a.DestApi); err != nil {
		return nil, err
	}
	if err = binary.Write(buf,binary.LittleEndian,a.DestAddr); err != nil {
		return nil, err
	}

	return buf,err
}

func (a *MessageB) Unmarshal(conn io.Reader) (*MessageB,error) {
	buf,err := _read(conn)
	if err != nil {
		return nil,err
	}
	var tmp = new(MessageB)
	tmp.MsgId = binary.LittleEndian.Uint32(buf[:4])
	tmp.Msg = buf[4:len(buf)-12]
	tmp.SrcAddr = binary.LittleEndian.Uint32(buf[len(buf)-12:len(buf)-12+4])
	tmp.DestApi = binary.LittleEndian.Uint32(buf[len(buf)-12+4:len(buf)-12+4+4])
	tmp.DestAddr = binary.LittleEndian.Uint32(buf[len(buf)-12+4+4:len(buf)-12+4+4+4])
	return tmp,nil
}

func (a *MessageB) Reply(msg []byte) (*bytes.Buffer,error) {
	a.Msg = msg
	return a.Marshal()
}