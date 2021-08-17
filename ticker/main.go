package main

import (
	"fmt"
	"time"
)

//打点器：固定时间间隔重复执行某个动作

func main() {
	ticker1 := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done: //等待一个完成消息
				return
			case t := <-ticker1.C:
				fmt.Println("ticker time:", t)
			}
		}
	}()

	//1600ms后停止
	time.Sleep(1600 * time.Millisecond)
	ticker1.Stop()
	done <- true

}
