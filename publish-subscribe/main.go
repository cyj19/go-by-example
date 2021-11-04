package main

import (
	"fmt"
	"github.com/vagaryer/go-by-exmaple/publish-subscribe/pubsub"
	"strings"
	"time"
)

// 有两个订阅者分别订阅全部主题和包含”golang“的主题
func main() {
	p := pubsub.NewPublisher(100*time.Microsecond, 10)
	defer p.Close()
	// 订阅全部主题
	all := p.Subscribe()
	// 订阅包含”golang“的主题
	golang := p.SubscribeTopic(func(v interface{}) bool {
		if s, ok := v.(string); ok {
			return strings.Contains(s, "golang")
		}
		return false
	})

	// 发布主题
	p.Publish("hello world")
	p.Publish("hello golang")

	go func() {
		for msg := range all {
			fmt.Println("全部主题", msg)
		}
	}()

	go func() {
		for msg := range golang {
			fmt.Println("golang主题", msg)
		}
	}()

	time.Sleep(3 * time.Second)
}
