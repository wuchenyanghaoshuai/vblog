package test

import (
	_ "github.com/wuchenyanghaoshuai/vblog/apps"
	"github.com/wuchenyanghaoshuai/vblog/conf"
	"github.com/wuchenyanghaoshuai/vblog/ioc"
)

func DevelopmentSetup() {
	err := conf.LoadConfigFromEnv()
	if err != nil {
		panic(err)
	}
	if err := ioc.Controller().Init(); err !=nil{
		panic(err)
	}
}
