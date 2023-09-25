package main

import (
	"fmt"


	"os"

	"github.com/gin-gonic/gin"
	tokenApiHandler "github.com/wuchenyanghaoshuai/vblog/apps/token/api"
	toeknImpl "github.com/wuchenyanghaoshuai/vblog/apps/token/impl"
	userImpl "github.com/wuchenyanghaoshuai/vblog/apps/user/impl"
	"github.com/wuchenyanghaoshuai/vblog/conf"
)

func main() {
	//1. 加载配置
	err := conf.LoadConfigFromToml("etc/application.toml")
	if err != nil {
		fmt.Println(err)	
		os.Exit(1)
	}
	//2. 初始化控制
	//user controller
	userServiceImpl :=userImpl.NewUserServiceImpl()
	//token controller
	tokenServiceImpl := toeknImpl.NewTokenServiceImpl(userServiceImpl)
	// token api handler
	tkApiHandler := tokenApiHandler.NewTokenApiHandler(tokenServiceImpl)

	//3.启动http协议服务器，注册handler路由

	r := gin.Default()
	tkApiHandler.Registry(r.Group("/api/vblog"))
	// 设置路由

	// 启动HTTP服务器
//	fmt.Println("http://127.0.0.1:7080")
	addr := conf.C().App.HttpAddr()
	fmt.Printf("HTTP API 监听地址: %s\n", addr)
	err = r.Run(addr)
	fmt.Println(err)

}
