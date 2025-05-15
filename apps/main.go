package main

import (
	"fmt"
	"os"
	//"vblog/apps/api"
	"vblog/apps/conf"
	"vblog/apps/ioc"

	//"github.com/gin-gonic/gin"
	//引用业务对象
	_ "vblog/apps/token/impl"
	_ "vblog/apps/user/impl"
	_ "vblog/apps/api"
)

func main() {
	//加载配置
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = "etc/application.yaml"
	}
	fmt.Println("configPath:", configPath)
	if err := conf.LoadConfigFromYaml(configPath)
		err != nil {
		panic(err)
	}


	//初始化ioc Controller
	err := ioc.Controller.Init()
	if err != nil {
		panic(err)	
	}
	//初始化ioc Api
	err = ioc.Api.Init()
	if err != nil {
		panic(err)	
	}
	// r := gin.Default()
	// root := r.Group("/vblog/api/v1") 
	// //把程序需要的对象组装起来，组装成业务程序，来处理业务逻辑
	// // userServiceImpl :=user.NewUserServiceImpl()
	// // tokenServiceImpl := token.NewTokenServiceImpl(userServiceImpl)


	// tokenApiHandler := api.NewTokenApiHandler()
	// tokenApiHandler.Registry(root.Group("tokens"))
	// r.Run(":8090")
	if err := conf.C().Application.Start(); err != nil {
		panic(err)
	}
}