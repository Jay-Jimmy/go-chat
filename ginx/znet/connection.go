package znet

import (
	"fmt"
	"go-chat/ginx/iface"
	"net"
)

type Connection struct {
	Conn     *net.TCPConn
	ConnID   uint32
	isClosed bool
	ExitChan chan bool
	Router   iface.IRouter
}

// 链接的读业务方法
func (c *Connection) StartReader() {
	fmt.Printf("Reader goroutine is running...")
	defer fmt.Printf("Reader exist")
	defer c.Stop()
	for {
		buf := make([]byte, 512)
		_, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Printf("Read buffer error:", err)
			continue
		}
		req := Request{
			conn: c,
			data: buf,
		}
		go func(request iface.IRequest) {
			c.Router.PreHandle(request)
			c.Router.Handle(request)
			c.Router.PostHandle(request)
		}(&req)
	}
}

// 启动链接
func (c *Connection) Start() {
	fmt.Printf("Connection %d is Starting", c.ConnID)
	go c.StartReader()
}

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
func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

// 发送数据给客户端
func (c *Connection) Send(data []byte) error {
	return nil
}

// 初始化链接模块的方法
func NewConnection(conn *net.TCPConn, connID uint32, router iface.IRouter) *Connection {
	c := &Connection{
		Conn:     conn,
		ConnID:   connID,
		isClosed: false,
		Router:   router,
		ExitChan: make(chan bool, 1),
	}
	return c
}
