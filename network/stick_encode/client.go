package main

import (
	"go-basic/network/stick_encode/proto"
	"log"
	"net"
	"strconv"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:30000")
	if err != nil {
		log.Println("[client] net.Dial failed, err", err)
		return
	}
	defer conn.Close()
	for i := 0; i < 100; i++ {
		//封包
		msg :=strconv.Itoa(i) + " hello server! " + time.Now().Format("15:04:05.000")
		bytes, err := proto.Encode(msg)
		if err != nil {
			log.Println("[client] proto.Encode failed, err:", err)
			break
		}
		conn.Write(bytes)
	}
}
