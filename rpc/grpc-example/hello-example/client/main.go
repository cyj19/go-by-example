/**
 * @Author: vagary
 * @Date: 2021/11/16 10:40
 */

// 客户端使用gRPC插件生成的方法来调用RPC
package main

import (
	"context"
	"fmt"
	"hello"
	"io"
	"log"
	"time"

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
	defer func() {
		_ = conn.Close()
	}()

	// 创建客户端
	c := hello.NewHelloClient(conn)

	// 调用方法
	req := &hello.HelloRequest{Name: "vagary"}
	resp, err := c.SayHello(context.Background(), req)
	if err != nil {
		grpclog.Fatalln(err)
	}

	fmt.Println(resp.Message)

	/***** 双向通信 ******/
	// 获取流对象
	stream, err := c.Channel(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	// 启动一个协程向服务端发送数据
	go func() {
		for {
			if err := stream.Send(&hello.StreamRequest{Data: "CYJ"}); err != nil {
				log.Fatal(err)
			}

			time.Sleep(time.Second)
		}

	}()

	// 循环接收服务端的响应
	for {
		reply, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}

		fmt.Println("双向通信，服务端的返回：", reply.GetData())
	}
}
