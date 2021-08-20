package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

/*
	标准库flag包-命令行参数解析
*/

// 获取命令行参数，没有定义命令行标签
func osArgsDemo() {
	if len(os.Args) > 0 {
		for index, arg := range os.Args {
			fmt.Printf("args[%d]:%s \n", index, arg)
		}
	}
}

// 自定义命令行标签(常用于指定端口，配置文件或其他资源)
func customDemo() {
	// 定义命令行标签，需要注意的是：已经定义的命令行标签不能再次定义
	// 参数1：命令行标签名称  参数2：默认值 参数3：描述
	port := flag.Int("port", 8080, "端口")
	mod := flag.String("mod", "dev", "运行环境")
	send := flag.Bool("send", true, "是否发送邮件")
	sendTime := flag.Duration("sendTime", 5*time.Second, "发送邮件间隔")
	// 解析命令行标签
	flag.Parse()
	fmt.Println("port:", *port)
	fmt.Println("mod:", *mod)
	fmt.Println("send:", *send)
	fmt.Println("sendTime:", *sendTime)
}

// 第二种定义命令行标签的方式
func custoDemo2() {
	var port int
	var mod string
	var send bool
	var sendTime time.Duration
	flag.IntVar(&port, "port", 8080, "端口")
	flag.StringVar(&mod, "mod", "dev", "运行环境")
	flag.BoolVar(&send, "send", true, "是否发送邮件")
	flag.DurationVar(&sendTime, "sendTime", 5*time.Second, "发送邮件间隔")

	// 解析命令行标签
	flag.Parse()
	fmt.Println("port:", port)
	fmt.Println("mod:", mod)
	fmt.Println("send:", send)
	fmt.Println("sendTime:", sendTime)
}

func main() {
	//osArgsDemo()
	//customDemo()
	custoDemo2()
}
