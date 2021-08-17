package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p *person) action() {
	fmt.Println("running...")
}

func main() {
	p := &person{name: "cyj", age: 25}
	fmt.Println(p.name, p.age)
	p.name = "zhangsan"
	fmt.Println(p)
	p.action()
}
