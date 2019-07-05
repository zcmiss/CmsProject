package config

import (
	"encoding/json"
	"fmt"
	"os"
)

//服务端配置
type AppConfig struct {
	AppName    string `json:"app_name"`
	Port       string `json:"port"`
	StaticPath string `json:"static_path"`
	Mode       string `json:"mode"`
}

var ServConfig AppConfig

//初始化服务器配置
func InitConfig() *AppConfig {
	file, err := os.Open("/config.json")
	if err != nil {
		fmt.Println("服务初始化失败...")
	}
	decoder := json.NewDecoder(file)
	conf := AppConfig{}
	err = decoder.Decode(&conf)
	if err != nil {
		fmt.Println("服务初始化失败...")
	}
	return &conf
}
