package config

import (
	"bytes"
	"io/ioutil"
	"log"

	"github.com/spf13/viper"
)

type Server struct {
	Name string
}

const configType = "yml"

var Vip *viper.Viper

func InitConfig() {
	Vip = viper.New()
	readConfig(Vip, "./config/config.yml")
}

func readConfig(vip *viper.Viper, filePath string) {
	vip.SetConfigType(configType)
	configBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal("读取配置文件失败，error:", err)
	}

	err = vip.ReadConfig(bytes.NewReader(configBytes))
	if err != nil {
		log.Fatal("viper 加载配置文件失败，error:", err)
	}
}
