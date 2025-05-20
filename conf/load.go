package conf

import (
	"os"

	"gopkg.in/yaml.v3"
	"github.com/caarlos0/env/v6"
)

//配置加载
// 将 file or env ------> Config 映射
// 全局的一份

var config *Config


func C()*Config {
	if config == nil {
		config = DefaultConfig()
	}
	return config
}

func(c *Config)ToYAML()string {
	out,_:=yaml.Marshal(c)	
	return string(out)
}


//加载配置
//把外部的配置读取到config全局变量中
func LoadConfigFromYaml(configPath string) error{
	content,err := os.ReadFile(configPath)
	if err != nil {
		return err
	}
	config = C()
	return yaml.Unmarshal(content,config)
}

//从环境变量中加载配置
func LoadConfigFromEnv() error{
	config = C()
	
	return env.Parse(config)
}