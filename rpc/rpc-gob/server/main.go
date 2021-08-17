package main

import (
	"log"
	"net/http"
	"net/rpc"
)

type Param struct {
	Width, Height int
}

type Rect struct{}

// 计算面积
func (r *Rect) Aear(p Param, ret *int) error {
	*ret = p.Width * p.Height
	return nil
}

// 计算周长
func (r *Rect) Perimeter(p Param, ret *int) error {
	*ret = (p.Width + p.Height) * 2
	return nil
}

func main() {
	// 1. 注册服务
	rect := new(Rect)
	// 把rect服务注册到rpc
	rpc.Register(rect)
	// 2. 服务处理绑定到http协议上
	rpc.HandleHTTP()
	// 监听服务
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalln("服务异常退出...")
	}
}
