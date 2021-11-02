package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
	使用原子操作和互斥锁可以高效地实现单例模式
*/

type singleton struct {
}

var (
	instance    *singleton
	initialized uint32     // 数字型标记位
	mu          sync.Mutex // 互斥锁
	once        = new(Once)
)

// Instance 有没有觉得内部的代码有点熟悉
func Instance() *singleton {
	if atomic.LoadUint32(&initialized) == 1 {
		return instance
	}
	mu.Lock()
	defer mu.Unlock()

	if instance == nil {
		defer atomic.StoreUint32(&initialized, 1)
		instance = &singleton{}
	}
	return instance

}

/*
	我们把Instance的公共部分抽象出来，这不正是标准库Sync.Once吗
*/

type Once struct {
	mu   sync.Mutex
	done uint32
}

func (o *Once) Do(f func()) {
	if atomic.LoadUint32(&o.done) == 1 {
		return
	}

	o.mu.Lock()
	defer o.mu.Unlock()
	if o.done == 0 {
		defer atomic.StoreUint32(&o.done, 1)
		f()
	}
}

func Instance2() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}

func main() {
	fmt.Printf("instance: %+v \n", Instance2())
}
