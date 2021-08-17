package main

import "fmt"

func judgeType(i interface{}) {
	switch t := i.(type) {
	case bool:
		fmt.Println("i is bool")
	case int:
		fmt.Println("i is int")
	case string:
		fmt.Println("i is string")
	default:
		fmt.Printf("unkown type: %T \n", t)
	}
}

func main() {
	i := 0
	switch {
	case i > 0:
		fmt.Println("i > 0")
	default:
		fmt.Println("i < 0")
	}

	//判断类型
	judgeType(10.0)
}
