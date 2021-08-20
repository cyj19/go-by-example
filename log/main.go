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

	有关FileMode的解释
	-rwxrwxrwx
	权限rwx分别对应4 2 1，相加的值为7

	第1位：文件属性，一般常用的是"-"，表示是普通文件；"d"表示是一个目录。

	第2～4位：文件所有者的权限rwx (可读/可写/可执行)。

	第5～7位：文件所属用户组的权限rwx (可读/可写/可执行)。

	第8～10位：其他人的权限rwx (可读/可写/可执行)。

	在golang中，可以使用os.FileMode(perm).String()来查看权限标识：

	os.FileMode(0777).String()    //返回 -rwxrwxrwx  代表文件所有者可读可写可执行，文件所属用户组可读可读可写可执行，其他人可读可读可写可执行

	os.FileMode(0666).String()   //返回 -rw-rw-rw- 代表文件所有者可读可写，文件所属用户组可读可写，其他人可读可写

	os.FileMode(0644).String()   //返回 -rw-r--r-- 代表文件所有者可读可写，文件所属用户组可读，其他人可读

*/

func main() {
	// 返回flag选项，时间、文件名、行号等
	fmt.Println(log.Flags())
	// 设置标准logger的标准输出
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	// 设置日志前缀
	log.SetPrefix("[vagaryer]")
	// 配置日志输出位置，参数3:0644  0代表八进制 后三位代表文件权限
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
