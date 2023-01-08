# go-jeans 是一个打包套接字，字节流的包，主要用来解决TCP传输中 [粘包](https://blog.csdn.net/weixin_41047704/article/details/85340311) 的问题

![golang](https://img.shields.io/badge/golang-v1.19-blue)
![simple](https://img.shields.io/badge/simple-extend-green)
![tcp-Pack](https://img.shields.io/badge/tcp-pack-yellowgreen)
![serve](https://img.shields.io/badge/network_transmission-pack-red)

* ### go-jeans 会把消息打包成基于 [Protobuff](https://zhuanlan.zhihu.com/p/401958878) 的字节流
* ### 基于两种消息结构传输 ([MessageA](#消息结构)、[MessageB](#消息结构))

## 消息结构
#### MessageA 由消息ID、消息组成 适用范围（个人见解）：客户端、服务端简单交互
```go
//由消息ID、消息组成 适用范围（个人见解）：客户端、服务端简单交互
type MessageA struct {
    state         protoimpl.MessageState
    sizeCache     protoimpl.SizeCache
    unknownFields protoimpl.UnknownFields
    
    MsgId uint32 `protobuf:"varint,1,opt,name=MsgId,proto3" json:"MsgId,omitempty"`
    Msg   []byte `protobuf:"bytes,2,opt,name=Msg,proto3" json:"Msg,omitempty"`
}
```
#### MessageB 由消息ID、消息、源地址（或标识）、目的接口、目的地址组成 适用范围（个人见解）：客户端和服务端、客户端请求服务端转发到指定客户端。
```go
//由消息ID、消息、源地址（或标识）、目的接口、目的地址组成 适用范围（个人见解）：客户端和服务端、客户端请求服务端转发到指定客户端。
type MessageB struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MsgId    uint32 `protobuf:"varint,1,opt,name=MsgId,proto3" json:"MsgId,omitempty"`
	Msg      []byte `protobuf:"bytes,2,opt,name=Msg,proto3" json:"Msg,omitempty"`
	SrcAddr  uint32 `protobuf:"varint,3,opt,name=SrcAddr,proto3" json:"SrcAddr,omitempty"`
	DestApi  uint32 `protobuf:"varint,4,opt,name=DestApi,proto3" json:"DestApi,omitempty"`
	DestAddr uint32 `protobuf:"varint,5,opt,name=DestAddr,proto3" json:"DestAddr,omitempty"`
}


```
## 使用教程
* ### 在项目中导入包
    go get -u github.com/Li-giegie/go_jeans
* ### 打包
#### 选择适合的消息结构（MessageA、MessageB） 打包提供三种方法
```go
//打包字节
//入参 打包的消息内容 
// 返回值 打包字节流、错误
// 打包后的字节流可以直接通过socket进行发送
PackA(msg []byte) ([]byte,error)

//打包字符串
PackA_String(msg string) ([]byte,error)

////打包Json对象
PackA_JSON(obj interface{}) ([]byte,error)
```

* ### 拆包
#### 根据打包的消息结构（MessageA、MessageB）进行拆包 

```go
//入参一个实现了io.Reader接口的对象 在go中一般情况下为socket的 connect对象
// 拆解MessageA 结构的包
UnpackA(conn io.Reader) (*MessageA,error)

// 拆解MessageA 结构的包
UnpackB(conn io.Reader) (*MessageB,error)

```