package main

import (
	"fmt"
	"viper-example/config"
)

func main() {
	var server config.Server
	config.InitConfig()
	fmt.Println(config.Vip.Get("server.name"))
	err := config.Vip.Unmarshal(&server)
	fmt.Println(server, err)
}
