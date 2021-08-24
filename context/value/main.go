package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
	context.WithValue(parent Context, key, val interface{}) Context
	仅对API和进程间传递请求域的数据使用上下文值，而不是使用它来传递可选参数给函数。
	1. key不能是内置类型，需要自定义，结构体或者内置类型的别名
*/

// 自定义key类型
type TraceCode string

var wg sync.WaitGroup

func main() {
	// 设置超时为50毫秒的上下文
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	ctx = context.WithValue(ctx, TraceCode("TRACE_CODE"), "123456789")
	wg.Add(1)
	go work(ctx)
	time.Sleep(5 * time.Second)
	// 通知子goroutine结束
	cancel()
	wg.Wait()
	fmt.Println("over")
}

func work(ctx context.Context) {
	defer wg.Done()
	// 从上下文中获取
	key := TraceCode("TRACE_CODE")
	traceCode, ok := ctx.Value(key).(string)
	if !ok {
		fmt.Println("invalid trace code")
		return
	}
LOOP:
	for {
		fmt.Println("trace code:", traceCode)
		// 假设正常连接数据库耗时10毫秒
		time.Sleep(10 * time.Millisecond)
		select {
		case <-ctx.Done():
			break LOOP
		}
	}

	fmt.Println("work done")
}
