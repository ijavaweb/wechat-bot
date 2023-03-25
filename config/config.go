package config

import (
	"os"
	"sync"
)

// Configuration 项目配置
type Configuration struct {
	// service apikey
	ApiKey string `json:"api_key"`
	// 自动通过好友
	AutoPass bool `json:"auto_pass"`
}

var config *Configuration
var once sync.Once

// LoadConfig 加载配置
func LoadConfig() *Configuration {
	once.Do(func() {
		config = &Configuration{}
		apiKey := os.Getenv("OPENAI_API_KEY")
		config.ApiKey = apiKey
		config.AutoPass = true
	})
	return config
}
