package api

import (
	"github.com/gin-gonic/gin"
	"github.com/wuchenyanghaoshuai/vblog/apps/token"
	"github.com/wuchenyanghaoshuai/vblog/ioc"
	"github.com/wuchenyanghaoshuai/vblog/response"
)

func init() {
	ioc.ApiHandler().Registry(&TokenApiHandler{})
}

func (t *TokenApiHandler) Name() string {
	return token.AppName
}
func (t *TokenApiHandler) Init() error {
	t.svc = ioc.Controller().Get(token.AppName).(token.Service)
	return nil
}

type TokenApiHandler struct {
	svc token.Service
}

func (h *TokenApiHandler) Registry(r gin.IRouter) {
	v1 := r.Group("v1")
	v1.POST("/tokens/", h.Login)
	v1.DELETE("/tokens", h.Logout)
}

func (h *TokenApiHandler) Login(c *gin.Context) {
	// 获取用户的请求参数,参数在body里面
	//一定要使用json
	req := token.NewLoginRequest()
	err := c.BindJSON(req)
	if err != nil {
		response.Failed(c, err)
		return
	}

	// 执行逻辑
	ins, err := h.svc.Login(c.Request.Context(), req)
	if err != nil {
		response.Failed(c, err)

		return
	}
	// 返回响应
	//c.JSON(http.StatusOK,ins)
	response.Success(c, ins)
}

func (h *TokenApiHandler) Logout(*gin.Context) {

}
