package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wuchenyanghaoshuai/vblog/apps/token"
)

func NewTokenApiHandler(tokenServiceImpl token.Service)*TokenApiHandler{
	return &TokenApiHandler{
		svc: tokenServiceImpl,
	}
}

type TokenApiHandler struct{
	svc  token.Service

}


func(h *TokenApiHandler) Registry(r gin.IRouter){
	v1 := r.Group("v1")
	v1.POST("/tokens/",h.Login)
	v1.DELETE("/tokens",h.Logout)
}

func (h *TokenApiHandler) Login(c *gin.Context){
	// 获取用户的请求参数,参数在body里面
	//一定要使用json
	req := token.NewLoginRequest()
	err := c.BindJSON(req)
	if err != nil{
		c.JSON(http.StatusBadRequest,err.Error())
		return
	}
	
	// 执行逻辑
	ins,err := h.svc.Login(c.Request.Context(),req)
	if err !=nil{
		c.JSON(http.StatusBadRequest,err.Error())
		return
	}
	// 返回响应
	c.JSON(http.StatusOK,ins)
	
}
func(h *TokenApiHandler) Logout(*gin.Context){

}