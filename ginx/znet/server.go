package znet

import (
	"fmt"
	"go-chat/ginx/iface"
	"net"
)

type Server struct {
	Name      string
	IPVersion string
	IP        string
	Port      int
	Router    iface.IRouter
}

func (s *Server) Start() {
	fmt.Printf("[Start] Server is starting...\n")
	go func() {
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Printf("Resolve %s addr error:", s.IPVersion, err)
			return
		}
		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Printf("Listen %s error:", s.IPVersion, err)
			return
		}
		fmt.Printf("[Running] Server %s starting succeess,listening at %s:%d\n", s.Name, s.IP, s.Port)
		var cid uint32
		cid = 0
		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Printf("Accept error:", err)
				continue
			}
			//实例化自定义链接
			dealConn := NewConnection(conn, cid, s.Router)
			cid++

			go dealConn.Start()
		}
	}()
}

func (s *Server) Stop() {

}

func (s *Server) Serve() {
	s.Start()
	select {}
}

func (s *Server) AddRouter(router iface.IRouter) {
	s.Router = router
	fmt.Printf("Add router succ\n")
}

func NewServer(name string) iface.IServer {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8999,
		Router:    nil,
	}
	return s
}
