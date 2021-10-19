package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:9000")
	if err != nil {
		log.Println("[server] net.Listen failed, err:", err)
		return
	}
	//阻塞等待客户端建立连接请求
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Println("[server] listen.Accept failed, err:", err)
			//失败继续阻塞
			continue
		}
		go handler(conn)
	}
}

func handler(conn net.Conn) {
	defer conn.Close() //函数return之前关闭连接
	for {
		bytes := make([]byte, 1024)
		length, err := conn.Read(bytes)
		if err != nil {
			log.Println("[server] read from client failed, err:", err)
			break
		}
		fmt.Println("[server] received data form client:",string(bytes[:length]))

		msg := "hello client! " + time.Now().Format("15:04:05.000")
		_, err = conn.Write([]byte(msg))
		if err != nil {
			log.Println("[server] send data to client failed, err:", err)
			continue
		}
	}
}
