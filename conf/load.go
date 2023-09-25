package conf

import (
	"github.com/BurntSushi/toml"
	"github.com/caarlos0/env/v6"
)

var (
	config *Config=DefaultConfig()
)

func C() *Config {
	if config == nil {
		panic("请提前加载配置")
	}
	return config
}


func LoadConfigFromToml(filepath string) error {
	_, err := toml.DecodeFile(filepath, config)
	if err != nil {
		return err
	}
	return nil
}

func LoadConfigFromEnv() error {
	return env.Parse(config)
}
