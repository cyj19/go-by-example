/**
 * @Author: vagary
 * @Date: 2021/11/17 16:50
 */

// 客户端1进行发布操作
package main

import (
	"context"
	pb "github.com/vagaryer/go-by-example/rpc/grpc-example/proto/pubsub"
	"google.golang.org/grpc"
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

	// 发布消息
	_, err = client.Public(context.Background(), &pb.String{Value: "golang: hello Go"})
	if err != nil {
		log.Fatal(err)
	}

	_, err = client.Public(context.Background(), &pb.String{Value: "docker: hello Docker"})
	if err != nil {
		log.Fatal(err)
	}

}
