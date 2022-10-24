package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	fmt.Printf("Client starting...")
	conn, err := net.Dial("tcp", "127.0.0.1:8999")
	if err != nil {
		fmt.Printf("Client start error:", err)
		return
	}
	for {
		_, err := conn.Write([]byte("Hello ginx!"))
		if err != nil {
			fmt.Printf("Write conn error:", err)
			return
		}
		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Printf("Read buffer error:", err)
			return
		}
		fmt.Printf("Server call back: %s, cnt = %d\n", buf, cnt)
		time.Sleep(1 * time.Second)
	}

}
