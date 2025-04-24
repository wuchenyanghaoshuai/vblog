package api

import (
	"vblog/apps/response"
	"vblog/apps/token"

	"github.com/gin-gonic/gin"
)

func NewTokenApiHandler(tokenServiceImpl token.Service) *TokenApiHandler {
	return &TokenApiHandler{
		token: tokenServiceImpl,
	}
}

type TokenApiHandler struct {
	token token.Service
}

func(h *TokenApiHandler)Registry(appRouter gin.IRouter){
	// r := gin.Default()
	// r.Group("api").Group("v1")
	appRouter.POST("login",h.Login)
}


func (h *TokenApiHandler) Login(c *gin.Context) {
	
	//1. 获取http请求
	req := token.NewIssueTokenRequest("", "")
	
	if err := c.BindJSON(req); err != nil {
		response.Failed(err, c)
		return
	}
	//2. 业务处理
	tk ,err := h.token.IssueToken(c.Request.Context(), req)
	if err != nil {
		response.Failed(err, c)}
	//3. 返回结果
	response.Success(tk, c)
}