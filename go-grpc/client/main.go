package main

import (
	"context"
	"fmt"
	proto "go-basic/go-grpc/api/v1"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial(":8888", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("client grpc.Dial err:%v", err)
	}
	defer conn.Close()
	client := proto.NewStreamApiClient(conn)
	resp, err := client.Add(context.Background(), &proto.Req{
		NumOne: 4,
		NumTwo: 10,
	})
	if err != nil {
		log.Fatalf("client.Add err:%v", err)
	}
	fmt.Printf("resp:%v", resp)

}
