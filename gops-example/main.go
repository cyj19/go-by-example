/**
 * @Author: cyj19
 * @Date: 2021/12/10 15:29
 */

// gops使用例子

package main

import (
	_ "expvar"
	"github.com/google/gops/agent"
	"log"
	"net/http"
)

func main() {
	// 创建并监听 gops agent，gops 命令会通过连接 agent 来读取进程信息
	// 若需要远程访问，可配置 agent.Options{Addr: "0.0.0.0:6060"}，否则默认仅允许本地访问gops
	if err := agent.Listen(agent.Options{}); err != nil {
		log.Fatalf("agent.Listen err: %v \n", err)
	}

	http.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		_, _ = w.Write([]byte("Go 语言编程之旅"))
	})

	log.Fatal(http.ListenAndServe(":6060", http.DefaultServeMux))
}
