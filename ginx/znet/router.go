package znet

import "go-chat/ginx/iface"

type BaseRouter struct {
}

func (br *BaseRouter) PreHandle(request iface.IRequest)  {}
func (br *BaseRouter) Handle(request iface.IRequest)     {}
func (br *BaseRouter) PostHandle(request iface.IRequest) {}
