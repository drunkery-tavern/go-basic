package main

import (
	"bufio"
	"fmt"
	"go-basic/network/stick_encode/proto"
	"io"
	"log"
	"net"
)

func main() {

	listen, err := net.Listen("tcp", "127.0.0.1:30000")
	if err != nil {
		log.Println("[server] net.Listen failed, err:", err)
		return
	}
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Println("[server] accept failed, err:", err)
			continue
		}
		go handler(conn)
	}
}

func handler(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		decode, err := proto.Decode(reader)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println("[server] read from client failed, err:", err)
			break
		}
		fmt.Println("[server] received client data：", decode)
	}
}
