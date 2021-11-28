/**
 * @Author: cyj19
 * @Date: 2021/11/10 10:40
 */

// websocket 客户端例子
package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
	"os"
	"os/signal"
)

// 从命令参数读取
var addr = flag.String("addr", "localhost:8888", "http server address")

func main() {
	flag.Parse()
	log.SetFlags(0)
	// 监听退出信号
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, os.Kill)
	// 构建连接地址
	u := url.URL{Scheme: "ws", Host: *addr, Path: "/echo"}
	log.Println("client1 connecting to ", u.String())
	// 拨号
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("客户端连接异常：", err.Error())
	}
	defer func() {
		_ = conn.Close()
	}()
	done := make(chan struct{})
	// 启动协程接收
	go receive(conn, done)
	_ = send(conn, "hello websocket")

	// 监听管道
	select {
	case <-done:
		return
	case <-interrupt:
		_ = send(conn, "hello websocket interrupt")
	}
}

// 接收服务端消息
func receive(conn *websocket.Conn, done chan<- struct{}) {
	defer func() {
		_ = conn.Close()
	}()
	for {
		mt, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("client1 read message error:", err.Error())
			break
		}
		fmt.Printf("messageType:%d message:%s \n", mt, string(msg))
	}
	done <- struct{}{}
}

// 发送消息
func send(conn *websocket.Conn, msg string) error {
	return conn.WriteMessage(websocket.BinaryMessage, []byte(msg))

}
