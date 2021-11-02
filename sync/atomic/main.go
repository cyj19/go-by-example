package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup

	var count uint64 = 0
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			for i := 1; i <= 1000; i++ {
				atomic.AddUint64(&count, 1)
			}
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println(count)
}
