package main

import (
	"fmt"
	"time"
)

func main() {
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)
	limiter := time.Tick(200 * time.Millisecond)

	for r := range requests {
		//通过接收操作来阻塞，到达200ms处理一次请求的效果
		<-limiter
		fmt.Println("request", r, time.Now())
	}

	burstyLimiter := make(chan time.Time, 3)

	go func() {
		//总体速率保持200ms处理一次请求
		// range channel 会自动接收，直到通道关闭
		for t := range time.Tick(200 * time.Millisecond) {
			//由于是缓冲通道，所以前三个不阻塞
			burstyLimiter <- t
		}

		//第二种写法，不过select一般用于处理多个不同channel
		// for {
		// 	select {
		// 	case t, exist := <-time.Tick(200 * time.Millisecond):
		// 		if !exist {
		// 			return
		// 		} else {
		// 			burstyLimiter <- t
		// 		}
		// 	}
		// }
	}()

	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	burstyRequests := make(chan int, 5)

	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	for r := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", r, time.Now())
	}

}
