package main

import "fmt"

func ping(w chan<- string, msg string) {
	w <- msg
}

func pong(r <-chan string, w chan<- string) {
	msg := <-r
	w <- msg
}

func main() {
	w := make(chan string)
	r := make(chan string)
	go ping(w, "golang")
	go pong(w, r)
	fmt.Println(<-r)
}
