package main

import (
	"fmt"
	"log"
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
	conn, err := jsonrpc.Dial("tcp", ":8080")
	if err != nil {
		log.Fatalf("连接失败：%v", err)
	}

	ar := new(ArithResponse)
	err = conn.Call("Arith.Multiply", ArithParam{5, 10}, ar)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("mutiply:", ar.Pro)

	err = conn.Call("Arith.Divide", ArithParam{5, 9}, ar)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("quo:%d rem:%d \n", ar.Quo, ar.Rem)
}
