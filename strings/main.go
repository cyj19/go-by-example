package main

import (
	"fmt"
	"strings"
)

func main() {
	var p = fmt.Println

	p("Contains: ", strings.Contains("text", "t"))
	p("Count: ", strings.Count("text", "t"))
	p("Join: ", strings.Join([]string{"a", "b"}, "-"))
	p("Repeat: ", strings.Repeat("adb", 4))
	p("Replace: ", strings.Replace("foo", "o", "0", -1))
	p("Replace: ", strings.Replace("foo", "o", "0", 1))
	p("Split: ", strings.Split("a-b-c-d", "-"))

}
