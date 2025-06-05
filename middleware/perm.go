package middleware

import (
	"vblog/apps/user"

	"github.com/gin-gonic/gin"
)

//这是一个有参数的中间件
//通过一个函数返回一个中间件
func Require(user.Role) func(c *gin.Context) {
	return func(c *gin.Context) {
		//判断用户当前的身份是不是跟角色匹配
	}
}