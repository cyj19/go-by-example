package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type ArithParam struct {
	A, B int
}

type ArithResponse struct {
	Pro int
	Quo int
	Rem int
}

func main() {
	// 连接服务端
	conn, err := net.Dial("tcp", ":8080")
	//conn, err := jsonrpc.Dial("tcp", ":8080")
	if err != nil {
		log.Fatalf("连接失败：%v", err)
	}

	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))

	ar := new(ArithResponse)
	err = client.Call("Arith.Multiply", ArithParam{5, 10}, ar)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("mutiply:", ar.Pro)

	err = client.Call("Arith.Divide", ArithParam{5, 9}, ar)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("quo:%d rem:%d \n", ar.Quo, ar.Rem)
}
