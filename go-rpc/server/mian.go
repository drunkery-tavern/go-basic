package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

type Server struct {
}

type Request struct {
	NumOne int
	NumTwo int
}

type Response struct {
	Num int
}

func (s *Server) Add(req Request, res *Response) error {
	log.Println("Add")
	time.Sleep(5 * time.Second)
	res.Num = req.NumTwo + req.NumOne
	return nil
}

func main() {
	err := rpc.Register(new(Server))
	if err != nil {
		return
	}
	rpc.HandleHTTP()
	listen, err := net.Listen("tcp", ":8888")
	if err != nil {
		return
	}
	err = http.Serve(listen, nil)
	if err != nil {
		return
	}
}
