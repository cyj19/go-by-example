/**
 * @Author: cyj19
 * @Date: 2021/12/8 11:21
 */

// pprof分析Mutex导致的阻塞例子

package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"sync"
)

func init() {
	runtime.SetMutexProfileFraction(1)
}

func main() {
	var m sync.Mutex
	datas := make(map[int]struct{})
	for i := 0; i < 999; i++ {
		go func(i int) {
			m.Lock()
			defer m.Unlock()
			datas[i] = struct{}{}
		}(i)
	}

	log.Fatal(http.ListenAndServe(":6061", nil))

}
