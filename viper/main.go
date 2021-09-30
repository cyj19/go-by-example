package main

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/vagaryer/viper-example/config"
)

/*
	问题：
	一开始配置文件只有Server时，
	使用Unmarshal将配置文件内容转为Server，发现转换不成功
	之后在外添加一层结构体Configuration，转换成功
	所以Unmarshal传入的结构体不能包含基本数据类型???
*/

var conf config.Configuration

func main() {
	firstMethod()
	//secondMethod()
}

func firstMethod() {
	config.FirstInitConfig()
	err := config.Vip.Unmarshal(&conf)
	fmt.Println(conf, err)
}

func secondMethod() {
	config.SecondInitConfig()
	err := viper.Unmarshal(&conf)
	fmt.Println(conf, err)
}
