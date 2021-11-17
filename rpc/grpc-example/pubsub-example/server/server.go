/**
 * @Author: vagary
 * @Date: 2021/11/17 16:18
 */

// gRPC的流特性构造一个发布和订阅系统
package main

import (
	"context"
	"github.com/docker/docker/pkg/pubsub"
	pb "github.com/vagaryer/go-by-example/rpc/grpc-example/proto/pubsub"
	"google.golang.org/grpc"
	"log"
	"net"
	"strings"
	"time"
)

type PubsubService struct {
	pub *pubsub.Publisher
}

// 实现PubsubServiceServer interface

func (p *PubsubService) Public(ctx context.Context, s *pb.String) (*pb.String, error) {
	p.pub.Publish(s.GetValue())
	return &pb.String{}, nil
}

func (p *PubsubService) Subscribe(s *pb.String, stream pb.PubsubService_SubscribeServer) error {
	ch := p.pub.SubscribeTopic(func(v interface{}) bool {
		if key, ok := v.(string); ok {
			if strings.HasPrefix(key, s.GetValue()) {
				return true
			}
		}
		return false
	})

	for val := range ch {
		// 推送给客户端
		if err := stream.Send(&pb.String{
			Value: val.(string),
		}); err != nil {
			return err
		}
	}

	return nil
}

func NewPubsubService() *PubsubService {
	return &PubsubService{
		pub: pubsub.NewPublisher(100*time.Microsecond, 10),
	}
}

const addr = "localhost:8889"

var _ pb.PubsubServiceServer = (*PubsubService)(nil)

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	// 创建gRPC对象
	gs := grpc.NewServer()
	// 注册服务
	pub := NewPubsubService()
	pb.RegisterPubsubServiceServer(gs, pub)
	// 启动服务
	log.Println("服务器启动")
	if err = gs.Serve(lis); err != nil {
		log.Fatal("服务器异常退出：", err)
	}
}
