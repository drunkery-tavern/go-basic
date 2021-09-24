package main

import (
	"fmt"
	"log"
	"net/rpc"
	"time"
)

type Request struct {
	NumOne int
	NumTwo int
}

type Response struct {
	Num int
}

func main() {
	request := Request{
		NumOne: 1,
		NumTwo: 2,
	}

	var resp Response
	client, err := rpc.DialHTTP("tcp", ":8888")
	if err != nil {
		log.Fatalf("err:%v", err)
	}

	//同步
	//err = client.Call("Server.Add", request, &resp)
	//if err != nil {
	//	log.Fatalf("err:%v", err)
	//}

	//异步
	call := client.Go("Server.Add", request, &resp, nil)
	for {
		select {
		case <-call.Done:
			fmt.Printf("Add result : %v", resp)
			return
		default:
			time.Sleep(1 * time.Second)
			fmt.Println("doing otherthings")
		}
	}
	//fmt.Printf("Add result : %v", resp)
}
