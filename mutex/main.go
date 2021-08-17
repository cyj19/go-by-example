package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	state := make(map[int]int)
	mutex := &sync.Mutex{}
	var readOps uint64
	var wirteOps uint64

	for i := 0; i < 100; i++ {
		go func() {
			for {
				total := 0
				key := rand.Intn(5)
				mutex.Lock()
				total += state[key]
				//fmt.Println(state[key])
				mutex.Unlock()
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	for i := 0; i < 10; i++ {
		go func() {
			for {
				key := rand.Intn(5)
				val := rand.Intn(100)
				mutex.Lock()
				state[key] = val
				mutex.Unlock()
				atomic.AddUint64(&wirteOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)

	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOpsFinal: ", readOpsFinal)
	wirteOpsFinal := atomic.LoadUint64(&wirteOps)
	fmt.Println("wirteOpsFinal: ", wirteOpsFinal)

	mutex.Lock()
	fmt.Println("state: ", state)
	mutex.Unlock()
}
