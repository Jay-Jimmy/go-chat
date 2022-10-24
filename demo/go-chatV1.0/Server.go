package main

import "go-chat/ginx/net"

func main() {
	s := net.NewServer("[ginx v1.0]")
	s.Serve()
}
