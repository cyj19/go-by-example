package main

import (
	"fmt"
	"time"
)

func worker(w int, jobs <-chan int, results chan<- string) {
	for j := range jobs {
		fmt.Println("worker ", w, "start job ", j)
		time.Sleep(1 * time.Second)
		fmt.Println("worker ", w, "finish job ", j)
		results <- fmt.Sprintf("worker %d finish job %d", w, j)
	}
}

func main() {
	num := 5
	jobs := make(chan int, num)
	results := make(chan string, num)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for i := 1; i <= num; i++ {
		//如果results是无缓冲通道的话，在worker中进行results的发送操作一直阻塞，导致jobs的发送操作也一同阻塞
		jobs <- i
		fmt.Println("send ", i)
	}
	close(jobs)
	//fmt.Println("send all jobs")
	//一直阻塞，因为results不关闭，一直等待接收
	// for r := range results {
	// 	fmt.Println(r)
	// }
	// close(results)

	for i := 1; i <= num; i++ {
		<-results
	}

}
