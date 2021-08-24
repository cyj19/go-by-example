package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

type respData struct {
	resp *http.Response
	err  error
}

func main() {
	// 创建一个超时上下文
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// 最后释放子goroutine的资源
	defer cancel()
	doCall(ctx)
	fmt.Println("shutdown...")
}

func doCall(ctx context.Context) {
	// 自定义客户端
	tr := &http.Transport{
		//如果DisableKeepAlives为真，会禁止不同HTTP请求之间TCP连接的重用
		DisableKeepAlives: true,
	}
	client := &http.Client{
		Transport: tr,
	}

	// 创建通知响应结果的channel
	respChan := make(chan respData)
	// 创建请求
	req, err := http.NewRequest("GET", "http://127.0.0.1:8080/", nil)
	if err != nil {
		fmt.Printf("create request fail: %v \n", err)
		return
	}
	// 创建一个携带上下文的请求
	req = req.WithContext(ctx)
	var wg sync.WaitGroup
	wg.Add(1)
	// 函数结束后开始等待，直到获取响应
	defer wg.Wait()
	// 在goroutine中发送请求
	go func() {
		// 发送请求
		resp, err := client.Do(req)
		// 把响应写入channel
		data := respData{
			resp: resp,
			err:  err,
		}
		respChan <- data
		wg.Done()
	}()

	// 监听IO操作
	select {
	case <-ctx.Done(): // 响应超时
		fmt.Println("request api timeout...")
		// 由于超时接收不到响应，这里要告诉WaitGroup任务已完成
		wg.Done()
	case data := <-respChan:
		if data.err != nil {
			fmt.Printf("response fail: %v \n", err)
			return
		}
		defer data.resp.Body.Close()
		// 处理响应的数据
		buf, err := ioutil.ReadAll(data.resp.Body)
		if err != nil {
			fmt.Printf("read response fail: %v \n", err)
			return
		}
		fmt.Println(string(buf))
	}

}
