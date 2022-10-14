package conf

import "fmt"

type WebConfiguration struct {
	RedisConfiguration RedisConfiguration
	DBConfiguration    DBConfiguration
	WebServerPort               string
}

func LoadConfiguration(confPath string) WebConfiguration {
	var res WebConfiguration
	fmt.Println("载入配置文件准备启动 WEB 服务")
	res.Port = ":80"
	return res
}
