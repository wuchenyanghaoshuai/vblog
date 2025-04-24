package main

import (
	"vblog/apps/api"
	token "vblog/apps/token/impl"
	user  "vblog/apps/user/impl"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	root := r.Group("/vblog/api/v1") 
	//把程序需要的对象组装起来，组装成业务程序，来处理业务逻辑
	userServiceImpl :=user.NewUserServiceImpl()
	tokenServiceImpl := token.NewTokenServiceImpl(userServiceImpl)


	tokenApiHandler := api.NewTokenApiHandler(tokenServiceImpl)
	tokenApiHandler.Registry(root.Group("tokens"))
	r.Run(":8090")
}