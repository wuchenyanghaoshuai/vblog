package start

import (

	"os"
	"vblog/conf"


	"github.com/spf13/cobra"
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