package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["k1"] = 1
	m["k2"] = 2
	m["k3"] = 3

	fmt.Printf("len = %d\n", len(m))
	//删除一个键值对
	delete(m, "k2")

	fmt.Println(m)
	//第二个参数判断key是否存在
	if v, ok := m["k1"]; ok {
		fmt.Printf("k1 value : %d \n", v)
	}

	for k, v := range m {
		fmt.Printf("key:%s  value:%d \n", k, v)
	}
}
