package rpc

import (
	"encoding/gob"
	"fmt"
	"net"
	"testing"
)

type User struct {
	Name string
	Age  int
}

// 定义用于服务端注册的函数
func QueryUser(uid int) (User, error) {
	// 简单模拟数据库
	users := make(map[int]User)
	users[0] = User{"one", 20}
	users[1] = User{"two", 21}
	users[2] = User{"three", 22}

	// 模拟查询用户
	if user, ok := users[uid]; ok {

		return user, nil
	}

	return User{}, fmt.Errorf("无此用户：%d", uid)
}

func TestClientServer(t *testing.T) {
	// 编码中有一个字段是interface{}时，要注册一下
	gob.Register(User{})
	// 创建服务端
	addr := "127.0.0.1:8080"
	rpcName := "QueryUser"
	server := NewServer(addr)
	// 注册函数
	server.Register(rpcName, QueryUser)
	// 另开协程，启动服务，不会阻塞
	go server.Run()

	// 创建客户端
	conn, _ := net.Dial("tcp", addr)
	client := NewClient(conn)
	// 声明函数原型
	var query func(int) (User, error)
	// RPC调用
	client.Call(rpcName, &query)
	// 调用函数query
	var user User
	user, err := query(1)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("user:", user)

}
