package main

import "fmt"

/*
	Q：slice[i:j] j代表什么？
	A：i表示元素下标，j不是下标也不是个数，j-1才代表元素下标；slice[i:j] 表示 从下标i开始到下标j-1，元素个数为j-i, cap为i开始到原slice的最后一个元素的个数

	Q: slice扩容后，原来slice[i:j]的存放的指针会改变吗?
	A：不会，扩容后slice指向了新的底层数组，所以原slice[i:j]和扩容后的slice无任何关系
*/

func main() {
	s := []int{1, 2, 3, 4, 5, 6}
	c := make([]int, 6)
	//len(s) != len(c) 无法复制
	num := copy(c, s)
	fmt.Printf("num: %d c:%v \n", num, c)

	fmt.Println(s[1:3], len(s[1:3]), cap(s[1:3])) //[2 3] len=2 cap=5
	fmt.Println(s[0:1], len(s[0:1]), cap(s[0:1])) // [1] len=1 cap=6
	fmt.Println(s[2:], len(s[2:]), cap(s[2:]))    //[3 4 5 6] len=4 cap=4
	fmt.Println(s[:3], len(s[:3]), cap(s[:3]))    // [1 2 3] len=3 cap=6
	fmt.Println(s[1:5], len(s[1:5]), cap(s[1:5])) // [2 3 4 5] len=4 cap=5

	// 证明扩容前后s[1:5]指向的底层数组的地址不变
	fmt.Printf("扩容前s地址：%p \n", &s[0])
	var s1 = s[1:5]
	fmt.Printf("扩容前s[1:5]地址：%p \n", &s1[0])
	s = append(s, 7, 8)
	fmt.Printf("扩容后地址：%p \n", &s[0])
	fmt.Printf("扩容后s[1:5]地址：%p \n", &s1[0])
	s1[0] = 9
	fmt.Println(s)
	var p *[5]int
	a := [5]int{1, 2, 3, 4, 5}
	p = &a

	fmt.Println(reverseArr(p))
}

/*
	模拟实现append函数
	1. 创建新slice
	2. 检测容量
	3. copy旧slice到新slice
	4. 添加元素
*/
func appendInt(x []int, y int) []int {
	// 新slice
	var z []int
	// 检测x容量是否足够
	zlen := len(x) + 1

	if zlen <= cap(x) {
		// 容量足够, 创建一个len=len+1的slice
		z = x[:zlen]

	} else {
		// 容量不足，进行扩容
		zcap := zlen
		// 新容量小于原长度的2倍，则新容量=2*原长度
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}

		// 根据zlen和zcap进行初始化
		z = make([]int, zlen, zcap)
		// 把x中的元素复制到z中
		copy(z, x)
	}

	z[len(x)] = y
	return z

}

// 重写reverse函数，使用数组指针代替slice
func reverseArr(p *[5]int) *[5]int {

	l := len(p)
	var arr [5]int
	for i := 0; i < l; i++ {
		arr[i] = p[l-1-i]
	}
	return &arr
}
