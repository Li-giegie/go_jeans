package go_jeans

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"errors"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"io"
	"sync"
)

var count uint32

var lock sync.Mutex

func PackA(msg []byte) ([]byte,error) {
	lock.Lock()
	count++
	var msgA = &MessageA{
		MsgId:         count,
		Msg:           msg,
	}
	lock.Unlock()
	return _Pack(msgA)
}

func PackA_String(msg string)([]byte,error)   {
	return PackA([]byte(msg))
}

func PackA_JSON(obj interface{})([]byte,error)   {
	msg,err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	return PackA(msg)
}

func PackB(msg []byte,srcAddr ,DestApi ,DestAddr uint32) ([]byte,error) {
	lock.Lock()
	count++
	var msgB = &MessageB{
		MsgId:         count,
		Msg:           msg,
		SrcAddr:       srcAddr,
		DestApi:       DestApi,
		DestAddr:      DestAddr,
	}
	lock.Unlock()
	return _Pack(msgB)
}

func PackB_String(msg string,srcAddr ,DestApi ,DestAddr uint32)([]byte,error)  {
	return PackB([]byte(msg),srcAddr,DestApi,DestAddr)
}

func PackB_JSON(obj interface{},srcAddr ,DestApi ,DestAddr uint32)([]byte,error)   {
	msg,err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}

	return PackB(msg,srcAddr,DestApi,DestAddr)
}

func _Pack(m protoreflect.ProtoMessage) ([]byte,error) {
	var buf = new(bytes.Buffer)

	var mbuf []byte

	var err error

	mbuf,err = proto.Marshal(m)

	if err != nil {
		return nil,newErr("Marshal proto err:",err)
	}

	if err = binary.Write(buf,binary.LittleEndian,uint32(len(mbuf)));err != nil {
		return nil, newErr("write msg len err :",err)
	}
	_,err = buf.Write(mbuf)

	return buf.Bytes(),nil
}

func _readBuf(conn io.Reader) ([]byte,error) {
	var packLen = make([]byte,4)
	var err error
	if _,err = io.ReadFull(conn,packLen);err != nil {
		return nil, newErr("read data err -1:",err)
	}

	var data = make([]byte,binary.LittleEndian.Uint32(packLen))
	if _,err = io.ReadFull(conn,data); err != nil {
		return nil, newErr("read data err -2:",err)
	}
	return data,nil
}

func UnpackA(conn io.Reader) (*MessageA,error) {

	var msgA = new(MessageA)

	packData,err := _readBuf(conn)

	if err != nil {
		return nil,err
	}

	return msgA,proto.Unmarshal(packData,msgA)
}

func UnpackB(conn io.Reader) (*MessageB,error) {

	var msgB = new(MessageB)

	packData,err := _readBuf(conn)

	if err != nil {
		return nil,err
	}

	return msgB,proto.Unmarshal(packData,msgB)
}

func newErr(textOrErr ...interface{}) error {
	var errBuf = new(bytes.Buffer)

	for _, i := range textOrErr {
		switch val := i.(type) {
		case error:
			errBuf.WriteString(val.Error())
		case string:
			errBuf.WriteString(val)
		}
	}

	return errors.New(errBuf.String())
}
