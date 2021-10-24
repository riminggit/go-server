package projectConfig

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var AppConfig = &Config{}

func InitConfig() {
	// 基础信息
	baseInfo := BaseInfo{env: "dev"}

	fmt.Println(baseInfo.env, "baseConfig.env")
	// 读取配置文件路径
	path := getAppPath() + "/config/" + baseInfo.env + ".ini"

	// 设置默认值
	// config := &Config{ServerConfig: ServerConfig{
	// 	ENV:           "pro",
	// 	HTTP_PORT:     8081,
	// 	READ_TIMEOUT:  60,
	// 	WRITE_TIMEOUT: 60,
	// }}

	// 映射，一切竟可以如此的简单。
	err := ini.MapTo(AppConfig, path)
	if err != nil {
		fmt.Println("ini文件映射出错啦，错误是:", err)
	}

	fmt.Println(AppConfig, "baseConfig")
}
