/**
 * @Author: vagary
 * @Date: 2021/11/10 10:41
 */

// websocket服务端例子
package main

import (
	"flag"
	"log"
	"net/http"
)

var addr = flag.String("addr", "localhost:8888", "http server address")

// HandlerFunc 类型
func echo(w http.ResponseWriter, r *http.Request) {
	log.Println("开始处理...")
	conn, err := wu.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade error:", err)
		return
	}
	c := &connection{ws: conn}
	// 每个连接启动一个协程处理
	go c.Writer()
}

func main() {

	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/echo", echo)
	log.Println("服务器启动...")
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("服务器异常退出")
	}
}
