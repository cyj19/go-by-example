/**
 * @Author: cyj19
 * @Date: 2022/4/26 14:01
 */

package main

import (
	"go/ast"
	"go/parser"
	"go/token"
)

var src = `
	package hello
	import "fmt"
	func PrintHello(){
		fmt.Println("Hello World")
	}
`

func main() {
	fset := token.NewFileSet()
	f, _ := parser.ParseFile(fset, "./hello/hello.go", src, 0)

	ast.Inspect(f, func(node ast.Node) bool {
		ast.Print(fset, node)
		return true
	})

	//ast.Print(fset, f)
}
