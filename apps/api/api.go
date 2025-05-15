package api

import (
	"vblog/apps/conf"
	"vblog/apps/ioc"
	"vblog/apps/response"
	"vblog/apps/token"

	"github.com/gin-gonic/gin"
)

// func NewTokenApiHandler() *TokenApiHandler {
// 	return &TokenApiHandler{
// 		token: ioc.Controller.Get(token.AppName).(token.Service),
// 	}
// }

func init() {
	ioc.Api.Registry(token.AppName, &TokenApiHandler{})
}

func (h *TokenApiHandler) Init() error {
	 h.token = ioc.Controller.Get(token.AppName).(token.Service)
	 subRouter := conf.C().Application.GinRootRouter().Group("tokens")
	 h.Registry(subRouter)
	return nil
}

type TokenApiHandler struct {
	token token.Service
}

func(h *TokenApiHandler)Registry(appRouter gin.IRouter){
	// r := gin.Default()
	// r.Group("api").Group("v1")
	appRouter.POST("/",h.Login)
	appRouter.DELETE("/",h.Logout)
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
	//
	c.SetCookie(token.COOKIE_TOKEN_KEY, tk.AccessToken, tk.RefreshTokenExpiredAt, "/", conf.C().Application.Domain, false, true)
	//3. 返回结果
	//返回数据 c.abort() 中断请求，不让请求继续执行后面的逻辑
	response.Success(tk, c)
}
func (h *TokenApiHandler) Logout(c *gin.Context) {
	ak,err := c.Cookie(token.COOKIE_TOKEN_KEY)
	if err != nil {
		response.Failed(err, c)
		return
	}
	refresh_token:=c.GetHeader(token.REFRESH_HEADER_KEY)
	//1. 获取http请求
	req := token.NewRevolkTokenRequest(ak, refresh_token)
	tk,err := h.token.RevolkToken(c.Request.Context(), req)
	if err != nil {
		response.Failed(err, c)
		return
	}
	//2. 业务处理
	response.Success(tk, c)

}