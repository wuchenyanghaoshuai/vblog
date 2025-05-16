package start

import (
	"fmt"
	"os"
//	"vblog/apps/cmd"
	"vblog/apps/conf"
	"vblog/apps/ioc"

	"github.com/spf13/cobra"
	_ "vblog/apps/token/impl"
	_ "vblog/apps/user/impl"
	_ "vblog/apps/api"
)
var Cmd = &cobra.Command{
	Use:   "start",
	Short: "Start vblog server",
	Run: func(cmd *cobra.Command, args []string) {
	// Do Stuff Here
	//加载配置
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = "etc/application.yaml"
	}
	cobra.CheckErr(conf.LoadConfigFromYaml(configPath))
	fmt.Println("configPath:", configPath)
	//初始化ioc Controller
	cobra.CheckErr(ioc.Controller.Init())
	//初始化ioc Api
	cobra.CheckErr(ioc.Api.Init())
	//启动
	cobra.CheckErr(conf.C().Application.Start())
	},
}

var (
	testParam string
)

func init() {
	Cmd.Flags().StringVarP(&testParam, "test", "t","test", "config flag")
}