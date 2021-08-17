package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string, 1)

	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "c1 done"
	}()

	select {
	case str := <-c1:
		fmt.Println(str)
	case <-time.After(1 * time.Second):
		fmt.Println("c1 timeout")
	}

	c2 := make(chan string, 1)

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "c2 done"
	}()

	select {
	case str := <-c2:
		fmt.Println(str)
	case <-time.After(3 * time.Second):
		fmt.Println("c2 timeout")
	}
}
