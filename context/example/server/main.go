package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	addr := "127.0.0.1:8080"
	http.HandleFunc("/", indexHandler)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("服务异常退出: %v \n", err)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// 随机出现慢响应
	number := rand.Intn(2)
	fmt.Println(number)
	if number == 0 {
		time.Sleep(10 * time.Second)
		fmt.Fprintf(w, "slow response")
		return
	}
	fmt.Fprintf(w, "quick response")

}
