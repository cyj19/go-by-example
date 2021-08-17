package main

import "fmt"

func main() {
	str := "abcdefg"
	//index：字符的位置  v：字符本身
	for index, v := range str {
		fmt.Printf("index: %d value:%d \n", index, v)
	}
}
