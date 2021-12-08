/**
 * @Author: cyj19
 * @Date: 2021/12/8 10:15
 */

// pprof分析简单例子

package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

var datas []string

func main() {
	go func() {
		for {
			log.Printf("len: %d \n", Add("pprof example"))
			time.Sleep(10 * time.Millisecond)
		}
	}()

	log.Fatal(http.ListenAndServe(":6060", nil))
}

func Add(str string) int {
	data := []byte(str)
	datas = append(datas, string(data))
	return len(datas)
}
