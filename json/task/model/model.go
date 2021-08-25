package model

type Server struct {
	Ip   string `ini:"ip"`
	Port int    `ini:"port"`
}

type Mysql struct {
	Username string  `ini:"username"`
	Passwd   string  `ini:"passwd"`
	Database string  `ini:"database"`
	Host     string  `ini:"host"`
	Port     int     `ini:"port"`
	Timeout  float64 `ini:"timeout"`
}

type Config struct {
	Server Server `ini:"server"`
	Mysql  Mysql  `ini:"mysql"`
}
