package main

import (
	"fmt"
	"json-task/config"
	"json-task/model"
)

/*
	task:
	1. 将结构体序列化为配置文件并保存
	2. 将配置文件反序列化为结构体
*/

// struct序列化到文件
func marshalDemo() {
	server := model.Server{
		Ip:   "127.0.0.1",
		Port: 8080,
	}
	mysql := model.Mysql{
		Username: "root",
		Passwd:   "123456",
		Database: "test",
		Host:     "127.0.0.1",
		Port:     3306,
		Timeout:  1.2,
	}
	conf := model.Config{
		Server: server,
		Mysql:  mysql,
	}
	fmt.Println(config.MarshalFile(conf))
}

// 配置文件反序列为struct
func unmarshalDemo() {
	conf := &model.Config{}
	fmt.Println(config.UnMarshalFile(conf))
	fmt.Println(conf)
}

func main() {
	//marshalDemo()
	unmarshalDemo()
}
