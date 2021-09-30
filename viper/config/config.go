package config

import (
	"bytes"
	"io/ioutil"
	"log"

	"github.com/spf13/viper"
)

type Configuration struct {
	Server Server
}

const (
	configType      = "yml"
	configPath      = "./config"
	defaultFilename = "config"
)

var Vip *viper.Viper

func FirstInitConfig() {
	Vip = viper.New()
	readConfig(Vip, configPath+"/"+defaultFilename+"."+configType)
}

func readConfig(vip *viper.Viper, filePath string) {
	vip.SetConfigType(configType)                 // 设置配置文件类型
	configBytes, err := ioutil.ReadFile(filePath) // 读取配置文件
	if err != nil {
		log.Fatal("读取配置文件失败，error:", err)
	}

	err = vip.ReadConfig(bytes.NewReader(configBytes)) // 将配置文件读到vip中
	if err != nil {
		log.Fatal("viper 加载配置文件失败，error:", err)
	}
}

func SecondInitConfig() {
	viper.SetConfigType(configType)      // 设置配置文件类型
	viper.SetConfigName(defaultFilename) // 设置配置文件名称
	viper.AddConfigPath(configPath)      // 设置配置文件所在目录
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("viper.ReadInConfig error:", err)
	}
}
