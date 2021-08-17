package rpc

import (
	"errors"
	"fmt"
	"log"
	"net"
	"sync"
	"testing"
)

func TestSession_ReadWriter(t *testing.T) {
	// 定义地址
	addr := "127.0.0.1:8080"
	// 定义测试数据
	mydata := "hello"

	wg := sync.WaitGroup{}
	// 等待两个任务
	wg.Add(2)

	// 创建服务端
	go func() {
		defer wg.Done()

		lis, err := net.Listen("tcp", addr)
		if err != nil {
			log.Fatal(err)
		}
		// 监听连接
		conn, err := lis.Accept()
		if err != nil {
			log.Fatal(err)
		}
		// 创建会话
		s := NewSession(conn)
		// 向连接中写数据
		err = s.Write([]byte(mydata))
		if err != nil {
			log.Fatal(err)
		}
	}()

	// 创建客户端
	go func() {
		defer wg.Done()

		conn, err := net.Dial("tcp", addr)
		if err != nil {
			log.Fatal(err)
		}
		// 创建会话
		s := NewSession(conn)
		// 从连接中读取数据
		data, err := s.Read()
		if err != nil {
			log.Fatal(err)
		}
		if string(data) != mydata {
			log.Fatal(errors.New("read fail"))
		}
		fmt.Println("data:", string(data))
	}()

	wg.Wait()
}
