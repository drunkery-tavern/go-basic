package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	client, err := net.Dial("tcp", "127.0.0.1:9000")
	if err != nil {
		log.Println("[client] net.Dial err:", err)
		return
	}
	defer client.Close()
	for {
		time.Sleep(1 * time.Second)
		msg := "hello server! " + time.Now().Format("15:04:05.000")
		_, err := client.Write([]byte(msg))
		if err != nil {
			log.Println("[client] write message err:", err)
			continue
		}
		var buf [1024]byte
		n, err := client.Read(buf[:])
		if err != nil {
			log.Println("[client] received data from server failed, err:", err)
			continue
		}
		fmt.Println("[client] received data form server:", string(buf[:n]))
	}
}
