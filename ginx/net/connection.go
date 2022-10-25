package net

import (
	"fmt"
	"go-chat/ginx/iface"
	"net"
)

type Connection struct {
	Conn      *net.TCPConn
	ConnID    uint32
	isClosed  bool
	handleAPI iface.HandleFunc
	ExitChan  chan bool
}

// 启动链接
func (c *Connection) Start()

// 停止链接
func (c *Connection) Stop() {
	fmt.Printf("Connection %d is Stopping", c.ConnID)
	if c.isClosed == true {
		return
	}
	c.isClosed = true
	err := c.Conn.Close()
	if err != nil {
		fmt.Printf("Connection closed error:", err)
	}
	close(c.ExitChan)
}

// 获取当前链接绑定的socket conn
func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

// 获取当前的链接ID
func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

// 获取远程客户端的TCP状态IP PORT
func (c *Connection) RemoteAddr() net.Addr

// 发送数据给客户端
func (c *Connection) Send(data []byte) error

// 初始化链接模块的方法
func NewConnection(conn *net.TCPConn, connID uint32, callback_api iface.HandleFunc) *Connection {
	c := &Connection{
		Conn:      conn,
		ConnID:    connID,
		isClosed:  false,
		handleAPI: callback_api,
		ExitChan:  make(chan bool, 1),
	}
	return c
}
