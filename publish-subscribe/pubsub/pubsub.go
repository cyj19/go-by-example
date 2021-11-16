/**
 * @Author: vagary
 * @Date: 2021/11/4 11:31
 */

// Package pubsub implements a simple multi-topic pub-sub library
// pubsub包是一个简单的多主题发布订阅库
// 基本思路：发布者和订阅者是多对多的关系，发布者需要维护自己的主题和订阅队列
// 主题：一个过滤函数
// 订阅：增加一个订阅队列和订阅的主题
// 发布：遍历所有订阅队列，使用主题处理消息，并发送到订阅队列中
package pubsub

import (
	"sync"
	"time"
)

type (
	subscriber chan interface{}         // 订阅队列为一个管道
	topicFunc  func(v interface{}) bool // 主题是一个过滤器
)

// Publisher 发布者对象
type Publisher struct {
	mu          sync.RWMutex             // 读写锁
	buffer      int                      // 订阅队列的缓存大小
	timeout     time.Duration            // 发布超时时间
	subscribers map[subscriber]topicFunc // 订阅者信息
}

// NewPublisher 构建一个发布者，可设置超时时间和订阅队列缓存大小
func NewPublisher(publishTimeout time.Duration, buffer int) *Publisher {
	return &Publisher{
		buffer:      buffer,
		timeout:     publishTimeout,
		subscribers: make(map[subscriber]topicFunc),
	}
}

// SubscribeTopic 添加一个新的订阅者，订阅过滤器筛选后的主题
func (p *Publisher) SubscribeTopic(topic topicFunc) chan interface{} {
	ch := make(chan interface{}, p.buffer)
	p.mu.RLock()
	p.subscribers[ch] = topic
	p.mu.RUnlock()
	return ch
}

// Subscribe 添加一个新的订阅者，订阅全部主题
func (p *Publisher) Subscribe() chan interface{} {
	return p.SubscribeTopic(nil)
}

// Evict 退出订阅
func (p *Publisher) Evict(sub chan interface{}) {
	p.mu.Lock()
	defer p.mu.Unlock()
	// 从订阅者集合中删除
	delete(p.subscribers, sub)
	// 关闭管道
	close(sub)
}

// 发送主题
func (p *Publisher) sendTopic(sub subscriber, topic topicFunc, v interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	if topic != nil && !topic(v) {
		return
	}

	select {
	case sub <- v:
	case <-time.After(p.timeout):
	}
}

// Publish 发布一个主题
func (p *Publisher) Publish(v interface{}) {
	p.mu.RLock()
	defer p.mu.RUnlock()
	var wg sync.WaitGroup
	for sub, topic := range p.subscribers {
		wg.Add(1)
		go p.sendTopic(sub, topic, v, &wg)
	}
	wg.Wait()
}

// Close 关闭发布者对象，同时关闭所有订阅者通道
func (p *Publisher) Close() {
	p.mu.RLock()
	defer p.mu.RUnlock()
	for sub := range p.subscribers {
		delete(p.subscribers, sub)
		close(sub)
	}
}
