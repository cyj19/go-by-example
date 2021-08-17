package main

import (
	"fmt"
	"time"
)

func main() {
	one := make(chan string)
	two := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		one <- "one"
	}()
	go func() {
		time.Sleep(time.Second)
		two <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case str1 := <-one:
			fmt.Println(str1)
		case str2 := <-two:
			fmt.Println(str2)
		}
	}

}
