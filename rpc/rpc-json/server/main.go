package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// 用于注册
type Arith struct {
}

// 运算参数
type ArithParam struct {
	A, B int
}

// 返回给客户端的结果
type ArithResponse struct {
	// 乘积
	Pro int
	// 商
	Quo int
	// 余
	Rem int
}

// 乘法
func (a *Arith) Multiply(p ArithParam, ar *ArithResponse) error {
	ar.Pro = p.A * p.B
	return nil
}

// 除法
func (a *Arith) Divide(p ArithParam, ar *ArithResponse) error {
	if p.B == 0 {
		return errors.New("除数不能为0")
	}

	ar.Quo = p.A / p.B
	ar.Rem = p.A % p.B
	return nil
}

func main() {

	// 1. 注册服务
	arith := new(Arith)
	rpc.Register(arith)
	// 2. 使用Listen创建服务端
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("服务创建异常：%v", err)
	}
	for {
		// Accept接收监听器l获取的连接，然后将每一个连接交给DefaultServer服务。Accept会阻塞"
		conn, err := lis.Accept()
		if err != nil {
			continue
		}
		// ServeConn在单个连接上执行DefaultServer。ServeConn会阻塞，服务该连接直到客户端挂起。调用者一般应另开线程调用本函数："go serveConn(conn)"。ServeConn在该连接使用JSON编解码格式。
		go func(conn net.Conn) {
			fmt.Println("json client")
			jsonrpc.ServeConn(conn)
		}(conn)
	}

}
