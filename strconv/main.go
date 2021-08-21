package main

import (
	"fmt"
	"strconv"
)

/*
	strconv包实现了基本数据类型与其字符串表示的转换
*/

// string转int
func atoi() {
	i, err := strconv.Atoi("10")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Type:%T Value:%d", i, i)
}

// int转string
func itoa() {
	s := strconv.Itoa(10)
	fmt.Printf("Type:%T Value:%s", s, s)
}

// string转bool
func parseBool() {
	b, err := strconv.ParseBool("false")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Type:%T Value:%t", b, b)
}

/*
	Parse类函数用于转换字符串为给定类型的值：
	ParseBool()、ParseFloat()、ParseInt()、ParseUint()。
*/

/*
	string 转 int
	参数1：要转换的字符串
	参数2：进制
	参数3：期望的接收类型，
	必须能无溢出赋值的整数类型，
	0、8、16、32、64 分别代表 int、int8、int16、int32、int64；
	例如期望接收int8，那么返回的结果可以无负担进行uint8(i)
*/
func parseInt() {
	i, err := strconv.ParseInt("15", 10, 64)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Type:%T Value:%d", i, i)
}

// string转uint
func parseUint() {
	i, err := strconv.ParseUint("15", 10, 0)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Type:%T Value:%d", i, i)
}

// string转float
func parseFloat() {
	f, err := strconv.ParseFloat("15.1", 64)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Type:%T Value:%f", f, f)
}

/*
	Format系列函数实现了将给定类型数据格式化为string类型数据的功能
*/

// bool转string
func formatBool() {
	s := strconv.FormatBool(true)
	fmt.Printf("Type:%T Value:%s", s, s)
}

// int转string
func formatInt() {
	s := strconv.FormatInt(50, 10)
	fmt.Printf("Type:%T Value:%s", s, s)
}

// uint转string
func formatUint() {
	s := strconv.FormatUint(50, 10)
	fmt.Printf("Type:%T Value:%s", s, s)
}

/*
	float转string
	参数2：表示格式
	参数3：精度控制或个数控制
	参数4：表示参数1的来源类型
	详见标准库文档https://studygolang.com/pkgdoc
*/
func formatFloat() {
	s := strconv.FormatFloat(3.1415, 'E', -1, 64)
	fmt.Printf("Type:%T Value:%s", s, s)
}

//判断一个字符是否可以打印, 和unicode.IsPrint一样，r必须是：字母（广义）、数字、标点、符号、ASCII空格。
func isPrint() {
	fmt.Println(strconv.IsPrint('a'), strconv.IsPrint('A'), strconv.IsPrint(1), strconv.IsPrint('.'), strconv.IsPrint('?'))
}

func main() {
	//atoi()
	//itoa()
	//parseBool()
	//parseInt()
	//parseUint()
	//parseFloat()
	//formatBool()
	//formatInt()
	//formatUint()
	//formatFloat()
	isPrint()
}
