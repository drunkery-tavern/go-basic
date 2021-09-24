package main

import (
	"context"
	"fmt"
	"go-basic/go-grpc/api/v1"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	proto.UnimplementedStreamApiServer
}

func (s *server) Add(ctx context.Context, req *proto.Req) (resp *proto.Resp, err error) {
	fmt.Println(req.GetNumOne(), req.GetNumTwo())
	return &proto.Resp{Result: req.GetNumOne() + req.GetNumTwo()}, nil
}

func main() {
	listen, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("net.Listen err:%v", err)
	}
	s := grpc.NewServer()
	proto.RegisterStreamApiServer(s, &server{})
	err = s.Serve(listen)
	if err != nil {
		log.Fatalf("run Serve err:%v", err)
	}

}
