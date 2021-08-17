package main

import "fmt"

/*
	func sum(s ...int) 表示变参，接收任意个int类型的参数
	sum(s...) 表示解序列，传入一个slice，然后用...解开为一个个
	用在数组中，是：= [...]int{1,2,3} 表示一个长度为指定元素个数的数组，是数组不是slice
*/

func sum(a, b int) int {
	return a + b
}

//变参函数
func variableParameter(s ...int) {
	fmt.Println(s)
}

func main() {
	fmt.Println(sum(1, 2))
	s := []int{1, 2, 3}
	variableParameter(s...)
}
