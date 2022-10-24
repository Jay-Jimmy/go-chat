package net

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
		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Printf("Accept error:", err)
				continue
			}
			go func() {
				for {
					buf := make([]byte, 512)
					cnt, err := conn.Read(buf)
					if err != nil {
						fmt.Printf("Receive buffer error:", err)
						continue
					}
					fmt.Printf("Receive client buf: %s, cnt=%d\n", buf, cnt)
					if _, err := conn.Write(buf[:cnt]); err != nil {
						fmt.Printf("Write back buffer error:", err)
						continue
					}
				}
			}()
		}
	}()
}

func (s *Server) Stop() {

}

func (s *Server) Serve() {
	s.Start()
	select {}
}

func NewServer(name string) iface.IServer {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8999,
	}
	return s
}
