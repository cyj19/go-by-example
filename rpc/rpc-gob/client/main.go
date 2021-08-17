package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Param struct {
	Width, Height int
}

func main() {
	// 连接服务
	conn, err := rpc.DialHTTP("tcp", ":8080")
	if err != nil {
		log.Fatalf("连接异常：%v", err)
	}

	// 面积
	ret := 0
	p := Param{5, 10}
	// 调用服务端的Aear方法
	err = conn.Call("Rect.Aear", p, &ret)
	if err != nil {
		log.Fatalf("调用服务失败：%v", err)
	}
	fmt.Println("aear:", ret)

	// 周长
	err = conn.Call("Rect.Perimeter", p, &ret)
	if err != nil {
		log.Fatalf("调用服务失败：%v", err)
	}
	fmt.Println("perimeter:", ret)
}
