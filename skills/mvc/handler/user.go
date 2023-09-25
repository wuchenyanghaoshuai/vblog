package handler

import (
	"github.com/gin-gonic/gin"
	"wuchenyanghaoshuai/vblog/skills/mvc/controller"
	"wuchenyanghaoshuai/vblog/skills/mvc/model"
)

// 实现createuser接口
func CreateUser(c *gin.Context) {

	// 从http协议读取用户请求的参数
	req := &model.User{}
	c.Bind(req)

	//调用controller来执行业务逻辑
	controller.CreateUser(req)
	c.JSON(0, req)
}
