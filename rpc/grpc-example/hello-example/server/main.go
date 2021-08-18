package main

import (
	"context"
	"fmt"
	"hello"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

const (
	Addr = "127.0.0.1:8080"
)

type helleService struct{}

var HelloService = helleService{}

func (h *helleService) SayHello(ctx context.Context, in *hello.HelloRequest) (*hello.HelloResponse, error) {
	resp := new(hello.HelloResponse)
	resp.Message = fmt.Sprintf("hello %s", in.Name)
	return resp, nil
}

func main() {
	lis, err := net.Listen("tcp", Addr)
	if err != nil {
		grpclog.Fatalf("failed to listen: %v \n", err)
	}

	// 实例化server
	s := grpc.NewServer()
	// 注册HelloServer
	hello.RegisterHelloServer(s, &HelloService)

	grpclog.Println("listen on " + Addr)
	s.Serve(lis)
}
