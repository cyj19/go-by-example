package main

import (
	"errors"
	"fmt"
	"time"
)

/*
	循环队列构建核心
	1. 队列判空 front == rear
	2. 队列判满 front == (rear+1)%max，队列剩余一个空闲单位时表示队列已满
	3. 先进先出，入队，rear向前移动；出队，front向前移动
	4. 队列最多只能存max-1个元素

	改善为线程安全
	思路：通过channal进行出队入队
*/

const (
	max = 10000
)

// 定义循环队列
type Circular struct {
	queue [max]int      // 队列
	front int           // 出队位置
	rear  int           // 入队位置
	lock  chan struct{} // 锁
}

// 创建队列
func Create() *Circular {
	return &Circular{
		queue: [max]int{},
		front: 0,
		rear:  0,
		lock:  make(chan struct{}, 1),
	}
}

// 判满
func (c *Circular) IsFull() bool {
	if c.front == (c.rear+1)%max {
		return true
	}

	return false
}

// 判空
func (c *Circular) IsEmpty() bool {
	if c.front == c.rear {
		return true
	}

	return false
}

// 入队
func (c *Circular) Add(value int) error {
	// 判满
	if c.IsFull() {
		return errors.New("the queue is full")
	}
	c.lock <- struct{}{}
	defer func() {
		<-c.lock
	}()
	c.queue[c.rear] = value
	// rear 往前移动
	c.rear = (c.rear + 1) % max
	return nil
}

// 出队
func (c *Circular) Remove() (int, error) {
	// 判空
	if c.IsEmpty() {
		return 0, errors.New("the queue is empty")
	}
	c.lock <- struct{}{}
	defer func() {
		<-c.lock
	}()
	value := c.queue[c.front]
	// front往前移动
	c.front = (c.front + 1) % max
	return value, nil
}

// 获取队列元素个数
func (c *Circular) Size() int {
	return (c.rear + max - c.front) % max
}

// 清空队列
func (c *Circular) Clear() {
	c.front = 0
	c.rear = 0
}

/*
	循环队列是否线程安全
	cir.Size = max-1 说明线程安全
*/
func nonThreadSafe(cir *Circular) {

	for n := 0; n < 10; n++ {
		go func() {
			for i := 1; i <= 1000; i++ {
				if err := cir.Add(i); err != nil {
					break
				}
			}

		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println(cir.Size())
}

func main() {
	cir := Create()
	nonThreadSafe(cir)
}
