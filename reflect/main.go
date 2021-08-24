package main

import (
	"fmt"
	"reflect"
)

/*
	reflect核心函数
	1. reflect.TypeOf() 获取静态类型
	2. reflect.ValueOf() 获取变量的值
	3. Value.Elem() 获取指针变量指向的元素
	4. Type.NumField() Value.NumField() 获取字段个数
	5. Type.Field(x) Value.Field(x) 获取相应的字段
	6. Kind() 获取字段具体类型
	7. Value.Set系列方法，反射修改属性值
	8. Value.NumMethod() Type.NumMethod() 获取方法个数
	9. Value.Method(x) 获取对应的方法
	10. Value.MethodByName(x)  获取对应方法
	11. Value.Call(xxx) 反射调用方法

	PS: 反射可以获取变量的所有属性，但反射只能修改公开的变量或调用公开的方法
*/

type A struct {
	Name string
}

type Example struct {
	Num   int
	Price float64
	Str   string
	Scl   []A
}

func (e Example) Print() {
	fmt.Println(e)
}

// 反射获取变量的所有属性
func getDemo() {
	example := Example{
		Num:   1,
		Price: 15.5,
		Str:   "example",
	}
	// 获取静态类型
	t := reflect.TypeOf(example)
	// kind()获取具体类型
	fmt.Printf("具体类型：%T \n", t.Kind())
	// 获取值
	v := reflect.ValueOf(example)
	fmt.Println("--------------字段--------------")
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Printf("%s: %v %v \n", f.Name, f.Type, val)
	}
	fmt.Println("--------------方法--------------")
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf("%s %v \n", m.Name, m.Type)
	}

}

// 反射修改变量的属性值
func changeDemo() {
	example := Example{
		Num:   1,
		Price: 15.5,
		Str:   "example",
	}
	fmt.Println(example)
	// 修改值要传入指针，获取指针指向的元素
	v := reflect.ValueOf(&example).Elem()
	for i := 0; i < v.NumField(); i++ {
		val := v.Field(i)
		switch val.Kind() {
		case reflect.Int:
			val.SetInt(2)
		case reflect.Float64:
			val.SetFloat(20.5)
		case reflect.String:
			val.SetString("example2")
		case reflect.Slice:
			a := make([]A, 0)
			a = append(a, A{"a"})
			val.Set(reflect.ValueOf(a))
			// 或者使用AppendSlice把一个slice追加到另一个slice
			// val.Set(reflect.AppendSlice(val, reflect.ValueOf(a)))

		}
	}
	// 反射调用方法
	m := v.MethodByName("print")
	var ages []reflect.Value
	m.Call(ages)
}

func main() {
	changeDemo()
}
