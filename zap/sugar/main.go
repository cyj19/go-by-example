package main

import (
	"net/http"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

/*
	zap自定义日志打印
	lumberjack进行日志切割归档
*/

var sugarLogger *zap.SugaredLogger

func main() {
	InitLogger()
	// 同步刷新所有缓存日志条目
	defer sugarLogger.Sync()
	for i := 0; i < 10000; i++ {
		simpleHttpGet("www.baidu.com")
		simpleHttpGet("https://www.baidu.com")
	}

}

func InitLogger() {
	// 创建编码器
	encoder := getEncoder()
	// 创建日志文件
	WriteSyncer := getLogWriter()

	core := zapcore.NewCore(encoder, WriteSyncer, zapcore.DebugLevel)
	logger := zap.New(core, zap.AddCaller())
	sugarLogger = logger.Sugar()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	// 日期格式化
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// 使用大写字母记录日志级别
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	// f, _ := os.OpenFile("./test.log", os.O_WRONLY|os.O_CREATE, 0744)
	// return zapcore.AddSync(f)
	// 使用lumberjack进行日志切割
	lumberjackLogger := &lumberjack.Logger{
		Filename:   "./test.log", // 日志文件位置
		MaxSize:    1,            // 单个日志文件最大大小
		MaxAge:     30,           // 保留旧日志文件的最大天数
		MaxBackups: 5,            // 保留旧日志文件的最大天数
		Compress:   false,        // 是否压缩旧日志文件
	}
	return zapcore.AddSync(lumberjackLogger)
}

func simpleHttpGet(url string) {
	sugarLogger.Debugf("Try to hit GET request for %s", url)
	resp, err := http.Get(url)
	if err != nil {
		sugarLogger.Errorf("Error fetching url %s: error=%v", url, err)
		return
	} else {
		sugarLogger.Infof("Success... statusCode=%s, url=%s", resp.Status, url)
		resp.Body.Close()
	}
}
