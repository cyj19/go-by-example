package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

type readOp struct {
	key  int
	resp chan int
}

type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func main() {

	var readOps uint64
	var writeOps uint64

	reads := make(chan readOp, 1)
	writes := make(chan writeOp, 1)

	//接收
	go func() {
		//协程独有
		state := make(map[int]int)
		for {
			select {
			case read, exist := <-reads:
				if !exist {
					return
				} else {
					read.resp <- state[read.key]
				}
			case write, exist := <-writes:
				if !exist {
					return
				} else {
					state[write.key] = write.val
					write.resp <- true
				}
			}
		}
	}()

	//发起读请求
	for r := 0; r < 100; r++ {
		go func() {
			for {
				read := readOp{
					key:  rand.Intn(5),
					resp: make(chan int),
				}
				reads <- read
				<-read.resp
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()

	}

	//发起写请求
	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool),
				}
				writes <- write
				<-write.resp
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()

	}

	time.Sleep(time.Second)

	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOpsFinal: ", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOpsFinal: ", writeOpsFinal)

}
