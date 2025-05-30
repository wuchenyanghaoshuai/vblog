package cmd

import (
	"fmt"
	initCmd "vblog/cmd/init"
	"vblog/cmd/start"
	"vblog/conf"
	"vblog/ioc"

	_ "vblog/apps" // Import all apps to ensure they are registered
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "vblog",
	Short: "Vblog Api Server",
	Run: func(cmd *cobra.Command, args []string) {


		if len(args) > 0 {
			if args[0] == "version" {
				// Do Stuff Here
				fmt.Println("version: v0.0.1")
			}else{
				cmd.Help()
			}
		}
	  // Do Stuff Here
  },
}
var (
	configPath string
)

func Execute() error{
	cobra.OnInitialize(func() {
		//加载配置
		cobra.CheckErr(conf.LoadConfigFromYaml(configPath))
		//初始化ioc Controller
		cobra.CheckErr(ioc.Controller.Init())
		//初始化ioc Api
		cobra.CheckErr(ioc.Api.Init())
	})
	return RootCmd.Execute()
}

func init() {
	//这里可以添加子命令
	RootCmd.PersistentFlags().StringVarP(&configPath, "config", "c", "etc/application.yaml", "the service config file")
	RootCmd.AddCommand(initCmd.Cmd)
	RootCmd.AddCommand(start.Cmd)
}