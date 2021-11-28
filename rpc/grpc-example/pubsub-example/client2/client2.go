/**
 * @Author: vagary
 * @Date: 2021/11/17 17:09
 */

// 客户端2订阅主题
package main

import (
	"context"
	"fmt"
	pb "github.com/cyj19/go-by-example/rpc/grpc-example/proto/pubsub"
	"google.golang.org/grpc"
	"io"
	"log"
)

const addr = "localhost:8889"

func main() {
	// 创建连接
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		_ = conn.Close()
	}()

	// 创建客户端
	client := pb.NewPubsubServiceClient(conn)

	// 订阅消息
	stream, err := client.Subscribe(context.Background(), &pb.String{Value: "golang"})
	if err != nil {
		log.Fatal(err)
	}

	for {
		reply, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		fmt.Println("消息：", reply.GetValue())
	}
}
