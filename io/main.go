package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// 文件方式操作终端
func terminalDemo() {
	var buf [16]byte
	os.Stdin.Read(buf[:])
	os.Stdout.WriteString(string(buf[:]))
}

// 打开文件
func openDemo() {
	f, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("open file err: %v \n", err)
		panic(err)
	}
	// 关闭文件
	defer f.Close()
}

func writeDemo() {
	// 创建新文件
	f, err := os.Create("./test.txt")
	if err != nil {
		fmt.Printf("create file err: %v \n", err)
		panic(err)
	}
	// 关闭文件
	defer f.Close()
	f.Write([]byte("write some to file"))
}

func readDemo() {
	// 打开文件
	f, err := os.Open("./test.txt")
	if err != nil {
		fmt.Printf("open file err: %v \n", err)
		panic(err)
	}
	defer f.Close()
	// 定义文件读取的字节数组
	var buf [128]byte
	var contain []byte
	for {
		n, err := f.Read(buf[:])
		// 结束返回EOF
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("read file err: %v \n", err)
			return
		}
		// 把读取到的字节存放到contain中
		contain = append(contain, buf[:n]...)
	}
	fmt.Println(string(contain))
}

func copyDemo() {
	// 创建新文件
	dstFile, err := os.Create("./test2.txt")
	if err != nil {
		fmt.Printf("create file err: %v \n", err)
		return
	}
	defer dstFile.Close()
	// 打开源文件
	srcFile, err := os.Open("./test.txt")
	if err != nil {
		fmt.Printf("open file err: %v \n", err)
		return
	}
	// 定义源文件接收数组
	var buf [128]byte
	for {
		n, err := srcFile.Read(buf[:])
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("read file err: %v \n", err)
			return
		}
		// 写入目标文件
		dstFile.Write(buf[:n])
	}
}

//使用bufio写文件
func wr() {
	// 创建新文件
	// 参数3：文件权限，r 4  w 2  x 1
	f, err := os.OpenFile("./test3.txt", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open file err: %v \n", err)
		return
	}
	defer f.Close()
	write := bufio.NewWriter(f)
	for i := 0; i < 5; i++ {
		write.WriteString("use bufio write something to file...\n")
	}

	// 刷新缓冲区，强制写出
	write.Flush()
}

//使用bufio读取文件
func re() {
	f, err := os.OpenFile("./test3.txt", os.O_RDONLY, 0644)
	if err != nil {
		fmt.Printf("open file err: %v \n", err)
		return
	}
	defer f.Close()
	read := bufio.NewReader(f)
	for {
		line, _, err := read.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("read file err: %v \n", err)
			return
		}
		fmt.Println(string(line))
	}
}

//使用ioutil写文件
func wr2() {
	err := ioutil.WriteFile("./test4.txt", []byte("use ioutil write file \nuse ioutil write file \n"), 0666)
	if err != nil {
		fmt.Printf("ioutil write file err: %v \n", err)
		return
	}
}

//使用ioutil读文件
func re2() {
	content, err := ioutil.ReadFile("./test4.txt")
	if err != nil {
		fmt.Printf("ioutil read file err: %v \n", err)
		return
	}
	fmt.Println(string(content))
}

//模拟linux平台的cat命令
func cat(r *bufio.Reader) {

	for {
		//line, _, err := r.ReadLine()
		buf, err := r.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("read file err: %v \n", err)
			return
		}
		//fmt.Println(string(line))
		fmt.Fprintf(os.Stdout, "%s", string(buf))
	}
}

func main() {
	//terminalDemo()
	//writeDemo()
	//readDemo()
	//copyDemo()
	//wr()
	//re()
	//wr2()
	//re2()

	// 解析命令行参数
	flag.Parse()
	// NArg返回解析flag之后剩余参数的个数
	if flag.NArg() == 0 {
		//没有参数，默认从标准输入读取
		cat(bufio.NewReader(os.Stdin))
	}
	// 依次读取指定文件并输出内容
	for i := 0; i < flag.NArg(); i++ {
		f, err := os.Open(flag.Arg(i))
		if err != nil {
			fmt.Fprintf(os.Stdout, "open file %s err: %v \n", flag.Arg(i), err)
			continue
		}
		cat(bufio.NewReader(f))
	}

}
