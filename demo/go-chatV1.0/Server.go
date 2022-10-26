package main

import (
	"fmt"
	"go-chat/ginx/iface"
	"go-chat/ginx/znet"
)

type PingRouter struct {
	znet.BaseRouter
}

func (this *PingRouter) PreHandle(request iface.IRequest) {
	fmt.Println("Call Router PreHandle")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("before ping...\n"))
	if err != nil {
		fmt.Println("callback before ping error")
	}
}

func (this *PingRouter) Handle(request iface.IRequest) {
	fmt.Println("Call Router Handle")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("ping...\n"))
	if err != nil {
		fmt.Println("callback ping error")
	}
}

func (this *PingRouter) PostHandle(request iface.IRequest) {
	fmt.Println("Call Router PostHandle")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("after ping...\n"))
	if err != nil {
		fmt.Println("callback post ping error")
	}
}

func main() {
	s := znet.NewServer("[ginx v1.0]")
	s.AddRouter(&PingRouter{})
	s.Serve()
}
