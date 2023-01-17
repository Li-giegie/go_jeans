package go_jeans

import (
	"google.golang.org/protobuf/proto"
	"io"
)

//func NewMsgA(msg []byte) *MessageA {
//	lock.Lock()
//	defer lock.Unlock()
//	count++
//	return &MessageA{
//		MsgId:         count,
//		Msg:           msg,
//	}
//
//}

//func NewMsgA_String(msg string) *MessageA  {
//	return NewMsgA([]byte(msg))
//}
//
//func NewMsgA_JSON(obj interface{}) (*MessageA,error)   {
//	msg,err := json.Marshal(obj)
//	if err != nil {
//		return nil, err
//	}
//	return NewMsgA(msg),nil
//}
//
//func NewMsgB(msg []byte,srcAddr ,DestApi ,DestAddr uint32) *MessageB {
//	lock.Lock()
//	defer lock.Unlock()
//	count++
//
//	return &MessageB{
//		MsgId:         count,
//		Msg:           msg,
//		SrcAddr:       srcAddr,
//		DestApi:       DestApi,
//		DestAddr:      DestAddr,
//	}
//
//}
//
//func NewMsgB_String(msg string,srcAddr ,DestApi ,DestAddr uint32) *MessageB {
//	return NewMsgB([]byte(msg),srcAddr,DestApi,DestAddr)
//}
//
//func NewMsgB_JSON(obj interface{},srcAddr ,DestApi ,DestAddr uint32)(*MessageB,error)   {
//	msg,err := json.Marshal(obj)
//	if err != nil {
//		return nil, err
//	}
//
//	return NewMsgB(msg,srcAddr,DestApi,DestAddr),nil
//}




func UnpackA_Proto(conn io.Reader) (*MessageA_Proto,error) {

	var msgA = new(MessageA_Proto)

	packData,err := _read(conn)

	if err != nil {
		return nil,err
	}

	return msgA,proto.Unmarshal(packData,msgA)
}

func UnpackB(conn io.Reader) (*MessageB_Proto,error) {

	var msgB = new(MessageB_Proto)

	packData,err := _read(conn)

	if err != nil {
		return nil,err
	}

	return msgB,proto.Unmarshal(packData,msgB)
}


