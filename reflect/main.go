package main

import (
	"fmt"
	"reflect"
)

type A struct {
}

type Example struct {
	Num   int
	Price float64
	Str   string
	Scl   []A
}

func main() {
	example := Example{}
	t := reflect.TypeOf(example)
	v := reflect.ValueOf(example)

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Printf("%s: %v %v \n", f.Name, f.Type, val)
	}
}
