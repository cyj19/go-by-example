/**
 * @Author: vagary
 * @Date: 2021/11/16 10:40
 */

// 服务端重新实现gRPC插件生成的接口，并注册到gRPC中
package main

import (
	"context"
	"fmt"
	"hello"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

const (
	Addr = "127.0.0.1:8080"
)

type HelleServiceImpl struct{}

// SayHello 重新实现gRPC生成的HelloServer Interface
func (h *HelleServiceImpl) SayHello(ctx context.Context, in *hello.HelloRequest) (*hello.HelloResponse, error) {
	resp := new(hello.HelloResponse)
	resp.Message = fmt.Sprintf("hello %s", in.Name)
	return resp, nil
}

func (h *HelleServiceImpl) Channel(server hello.Hello_ChannelServer) error {
	for {
		// 接收客户端请求
		args, err := server.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		fmt.Println("双向通信，客户端请求：", args.GetData())
		reply := &hello.StreamResponse{Data: "stream " + args.GetData()}
		// 回复
		err = server.Send(reply)
		if err != nil {
			return err
		}
	}
}

var _ hello.HelloServer = (*HelleServiceImpl)(nil)

func main() {
	lis, err := net.Listen("tcp", Addr)
	if err != nil {
		grpclog.Fatalf("failed to listen: %v \n", err)
	}

	// 构建一个gRPC服务对象
	s := grpc.NewServer()
	// 注册HelloServer
	hello.RegisterHelloServer(s, new(HelleServiceImpl))

	grpclog.Info("listen on", Addr)
	if err = s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
