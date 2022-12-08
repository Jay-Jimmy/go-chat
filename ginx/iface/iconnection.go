package iface

import "net"

type IConnection interface {
	//启动链接
	Start()
	//停止链接
	Stop()
	//获取当前链接绑定的socket conn
	GetTCPConnection() *net.TCPConn
	//获取当前的链接ID
	GetConnID() uint32
	//获取远程客户端的TCP状态IP PORT
	RemoteAddr() net.Addr
	//发送数据给客户端
	Send(data []byte) error
}

// 定义一个处理链接业务的方法
type HandleFunc func(*net.TCPConn, []byte, int) error
