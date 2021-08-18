package main

import (
	"context"
	"fmt"
	"hello"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

const (
	Addr = "127.0.0.1:8080"
)

func main() {
	// 连接服务端
	conn, err := grpc.Dial(Addr, grpc.WithInsecure())
	if err != nil {
		grpclog.Fatalln(err)
	}
	defer conn.Close()

	// 创建客户端
	c := hello.NewHelloClient(conn)

	// 调用方法
	req := &hello.HelloRequest{Name: "chen"}
	resp, err := c.SayHello(context.Background(), req)
	if err != nil {
		grpclog.Fatalln(err)
	}

	fmt.Println(resp.Message)
}
