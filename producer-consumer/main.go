package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

/*
	生产者消费者模型
*/

// Producer 生成factor的整数倍
func Producer(factor int, out chan<- int) {
	for i := 0; ; i++ {
		out <- i * factor
	}
}

// Consumer 消费成果队列的数据
func Consumer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func main() {
	ch := make(chan int, 64) // 成果队列
	go Producer(3, ch)       // 生成3的倍数的序列
	go Producer(5, ch)       // 生成5的倍数的序列
	go Consumer(ch)          // 消费 生成的队列

	// Ctrl + C退出
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	fmt.Printf("quit (%v)\n", <-sig)
	os.Exit(0)
}
