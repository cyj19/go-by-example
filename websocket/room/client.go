/**
 * @Author: cyj19
 * @Date: 2021/12/1 15:09
 */

// TCP聊天室客户端
package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", ":8888")
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan struct{})

	go func() {
		// 从conn读取数据并写到标准输出
		_, _ = io.Copy(os.Stdout, conn)
		log.Println("done")
		done <- struct{}{}
	}()

	// 从标准输入获取输入数据
	mustCopy(conn, os.Stdin)

	<-done
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
