package main

import (
	"net/http"

	"go.uber.org/zap"
)

/*
	zap--高性能日志库
	优势：
		1. 可以打印不同级别的日志 (标准库log包只有Print系列打印，Fatal系列打印后会执行os.Exit(1)退出程序，Panic打印后会抛出异常)
		2. 可以打印被调用者的函数名
*/

var logger *zap.Logger

func main() {
	InitLogger()
	defer logger.Sync()
	simpleHttp("www.baidu.com")
	simpleHttp("https://www.baidu.com")
}

func InitLogger() {
	logger, _ = zap.NewProduction()
}

func simpleHttp(url string) {
	resp, err := http.Get(url)
	if err != nil {
		logger.Error("Error fetch url...", zap.String("url", url), zap.Error(err))
		return
	} else {
		logger.Info("Success...", zap.String("StatusCode", resp.Status), zap.String("url", url))
		resp.Body.Close()
	}
}
