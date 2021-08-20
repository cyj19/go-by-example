package main

import (
	"fmt"
	"log"
	"os"
)

/*
	标准库log包
	Print系列函数：标准错误输出
	Fanta系列函数：标准错误输出，输出后调用os.Exit(1)
	Panic系列函数：标准错误输出，输出后调用panic
*/

func main() {
	// 返回flag选项，时间、文件名、行号等
	fmt.Println(log.Flags())
	// 设置标准logger的标准输出
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	// 设置日志前缀
	log.SetPrefix("[vagaryer]")
	// 配置日志输出位置
	file, err := os.OpenFile("./log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err)
	}
	log.SetOutput(file)
	// 打印日志
	log.Println("测试日志")

	// 创建logger
	file2, err := os.OpenFile("./log2.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err)
	}
	mylog := log.New(file2, "<vagaryer>", log.Llongfile|log.Lmicroseconds|log.Ldate)
	mylog.Println("自定义创建的日志对象")
}
