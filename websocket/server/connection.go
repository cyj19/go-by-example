/**
 * @Author: vagary
 * @Date: 2021/11/10 11:24
 */

package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

// 使用自定义参数
var wu = &websocket.Upgrader{
	ReadBufferSize:   1024,
	WriteBufferSize:  1024,
	HandshakeTimeout: 5 * time.Second,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// 连接信息
type connection struct {
	ws *websocket.Conn
}

func (c *connection) Writer() {
	defer func() {
		_ = c.ws.Close()
	}()
	for {
		mt, msg, err := c.ws.ReadMessage()
		if err != nil {
			log.Println("server read message error:", err.Error())
			break
		}
		var msgStr = string(msg)
		fmt.Printf("messageType:%d, message:%s \n", mt, msgStr)
		// 回复
		_ = c.ws.WriteMessage(mt, []byte(msgStr+" server"))
	}
}
