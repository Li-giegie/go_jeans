package go_jeans

import (
	"fmt"
	"testing"
)

func TestB(t *testing.T) {
	msg:=NewMsgB([]byte("hello word B"),1,2,3)
	buf,err := msg.Marshal()
	fmt.Println(err)

	msgB,err := msg.Unmarshal(buf)
	fmt.Println(err)
	fmt.Println(string(msgB.Msg),msgB.MsgId)
}

func BenchmarkB(b *testing.B) {
	for i := 0; i < b.N; i++ {

		buf,_ := NewMsgB([]byte("hello word B"),1,2,3).Marshal()
		//fmt.Println(err)

		(&MessageB{}).Unmarshal(buf)
		//fmt.Println(err)
		//fmt.Println(string(msgB.Msg),msgB.MsgId)
	}
}

func BenchmarkName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		
	}
}