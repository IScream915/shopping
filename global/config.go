package global

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

const DefaultGlobalConfigPath = "config/globalConfig.yaml"

type Config struct {
	Salt *Salt `mapstructure:"salt"`
}

func NewConfig() *Config {
	// 读取文件内容
	config, err := os.ReadFile(DefaultGlobalConfigPath)
	if err != nil {
		log.Fatalln("读取globalConfig失败: ", err)
	}

	// 反序列化到结构体
	cfg := &Config{}
	if err := yaml.Unmarshal(config, cfg); err != nil {
		log.Fatalln("解析 yaml 失败: ", err)
	}

	return cfg
}
