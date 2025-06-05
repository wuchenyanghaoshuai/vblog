package middleware

import (
	"vblog/apps/token"
	"vblog/ioc"
	"vblog/response"
	"github.com/gin-gonic/gin"
)

//Gin Web 中间件,我们需要在中间件注入到请求的链路中，然后由Gin框架来调用
func Auth(c *gin.Context) {
	// 1. 获取请求头中的token
	accessToken,_ := c.Cookie(token.COOKIE_TOKEN_KEY)
	tk,err := ioc.Controller.Get(token.AppName).(token.Service).ValidateToken(c.Request.Context(),token.NewValidateTokenRequest(accessToken))
	if err != nil {
		//响应报错信息
		response.Failed(token.ErrAuthFailed.WithMessage(err.Error()),c)
		c.Abort()
		
	}else{
		//后面的handler怎么知道是谁在访问这个接口，以及鉴权成功了
		//请求上下文机制
		c.Set(token.GIN_TOKEN_KEY_NAME,tk)
		c.Next()

	}
	
	
}