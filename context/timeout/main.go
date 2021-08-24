package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
	context.WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
	当超过指定事件，Done通道会关闭
*/

var wg sync.WaitGroup

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	wg.Add(1)
	go work(ctx)
	time.Sleep(5 * time.Second)
	// 执行完，立即释放资源
	cancel()
	wg.Wait()
	fmt.Println("over")

}

// 模拟数据库连接
func work(ctx context.Context) {
	defer wg.Done()
LOOP:
	for {
		fmt.Println("db connecting...")
		// 假设数据库连接
		time.Sleep(10 * time.Millisecond)
		select {
		case <-ctx.Done(): // 50毫秒后自动调用
			break LOOP
		}
	}
}
