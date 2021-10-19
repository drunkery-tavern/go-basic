package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("[client] net.Dial failed, err", err)
		return
	}
	defer conn.Close()
	for i := 0; i < 100; i++ {
		msg := "hello server! " + time.Now().Format("15:04:05.000")
		conn.Write([]byte(msg))
	}
}
