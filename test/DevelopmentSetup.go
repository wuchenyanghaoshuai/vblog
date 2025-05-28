package test

import (
	_ "vblog/apps"
	"vblog/conf"
	"vblog/ioc"

	"github.com/spf13/cobra"
)

func DevelopmentSetup() {
	// This function is a placeholder for setting up the development environment.
	// It can be used to initialize configurations, set up databases, or any other
	// necessary setup before running tests.
	if err := conf.LoadConfigFromEnv(); err != nil {
		panic(err)
	}
	 	//初始化ioc Controller
	cobra.CheckErr(ioc.Controller.Init())
	//初始化ioc Api
	cobra.CheckErr(ioc.Api.Init())
}