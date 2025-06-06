package middleware

import (
	"vblog/apps/token"
	"vblog/apps/user"
	"vblog/response"

	"github.com/gin-gonic/gin"
)

//这是一个有参数的中间件
//通过一个函数返回一个中间件
//这个中间件是加载在认证中间件之后的
func RequireRole(requiredRoles ...user.Role) func(c *gin.Context) {
	return func(c *gin.Context) {
		//判断用户当前的身份是不是跟角色匹配
		if v,ok := c.Get(token.GIN_TOKEN_KEY_NAME);ok{
			for i := range requiredRoles{
				requiredRole := requiredRoles[i]
				if v.(*token.Token).Role == requiredRole{
					c.Next()
				}
			}
		}
		response.Failed(token.ErrPermissionDenied,c)
		c.Abort()
	}
}