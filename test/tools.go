package test

import "github.com/wuchenyanghaoshuai/vblog/conf"

func DevelopmentSetup() {
	err := conf.LoadConfigFromEnv()
	if err != nil {
		panic(err)
	}
}
