package go_jeans

import (
	"bytes"
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPackA_go(t *testing.T) {
	//fmt.Println(2)
	var s sync.WaitGroup
	var n = 1000000
	var msgids = make([]uint32,n,n)
	t1 := time.Now()
	for i:=0;i<n;i++{
		s.Add(1)
		go func(j int) {
			defer s.Done()
			data,err := PackA([]byte("hello word"))
			if err != nil {
				t.Error(err)
			}
			var buf = bytes.NewBuffer(data)
			msgA,err := UnpackA(buf)
			if err != nil {
				t.Error(err,msgA)
			}
			//fmt.Println("write1 ",j,msgA.MsgId)
			msgids[j] = msgA.MsgId
			//fmt.Println("write2 ",msgA.MsgId,msgids[j])
		}(i)
	}
	s.Wait()
	fmt.Println("打包完毕 耗时",time.Since(t1))
	tmp := 0
	//fmt.Println(len(msgids),msgids)
	for i, msgid := range msgids {
		tmp = 0
		for _, u := range msgids {
			if msgid == u {
				tmp++
			}
		}
		if tmp > 1 {
			fmt.Println("重复id：",msgid,tmp)
		}
		if (len(msgids) - i) % 10000 == 0 {
			fmt.Println(len(msgids)-i)
		}
	}

}

func TestPackA(t *testing.T) {

	var n = 1000000
	var msgids = make([]uint32,n,n)
	t1 := time.Now()

	for i:=0;i<n;i++{
		data,err := PackA([]byte("hello word"))
		if err != nil {
			t.Error(err)
		}
		var buf = bytes.NewBuffer(data)
		msgA,err := UnpackA(buf)
		if err != nil {
			t.Error(err,msgA)
		}
		msgids[i] = msgA.MsgId

	}

	fmt.Println("打包完毕 耗时",time.Since(t1))
	tmp := 0
	//fmt.Println(len(msgids),msgids)
	for i, msgid := range msgids {
		tmp = 0
		for _, u := range msgids {
			if msgid == u {
				tmp++
			}
		}
		if tmp > 1 {
			fmt.Println("重复id：",msgid,tmp)
		}
		if (len(msgids) - i) % 10000 == 0 {
			fmt.Println(len(msgids)-i)
		}
	}

}

func Benchmark_PackA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		data,err := PackA([]byte("hello word"))
		if err != nil {
			b.Error(err)
		}

		//var buf = bytes.NewBuffer([]byte("asdasda&&&^^^"))
		var buf = bytes.NewBuffer(data)
		//buf.Write(data)

		msgA,err := UnpackA(buf)
		if err != nil {
			b.Error(err,msgA)
		}
	}
}