package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/wuchenyanghaoshuai/vblog/apps/token"
	"github.com/wuchenyanghaoshuai/vblog/apps/user"
	"github.com/wuchenyanghaoshuai/vblog/exception"
	"github.com/wuchenyanghaoshuai/vblog/ioc"
	"github.com/wuchenyanghaoshuai/vblog/response"
	"net/http"
)

//用于鉴权的中间件

func NewTokenAuther() *TokenAuther {
	return &TokenAuther{
		tk: ioc.Controller().Get(token.AppName).(token.Service),
	}
}

type TokenAuther struct {
	tk   token.Service
	role user.Role
}

// 怎么鉴权
// Gin中间件 func(*Context)
func (a *TokenAuther) Auth(c *gin.Context) {
	//1.获取token
	at, err := c.Cookie(token.TOKEN_COOKIE_NAME)
	if err != nil {
		if err == http.ErrNoCookie {
			response.Failed(c, token.CookieNotFound)
			return
		}
		response.Failed(c, err)
		return
	}
	//2.调用token模块来验证
	in := token.NewValidateToken(at)
	tk, err := a.tk.ValidateToken(c.Request.Context(), in)
	if err != nil {
		response.Failed(c, err)
		return
	}
	//把鉴权后的结果：tk 放到请求的上下文，方便后面的业务逻辑来使用
	if c.Keys == nil {
		c.Keys = map[string]any{}
	}
	c.Keys[token.TOKEN_GIN_KEY_NAME] = tk
	//	return
}

// 权限的鉴定
// 判断当前用户的角色
func (a *TokenAuther) Perm(c *gin.Context) {
	tkObj := c.Keys[token.TOKEN_GIN_KEY_NAME]
	if tkObj == nil {
		response.Failed(c, exception.NewPermissionDeny("tooken not found"))
		return
	}
	tk, ok := tkObj.(*token.Token)
	if !ok {
		response.Failed(c, exception.NewPermissionDeny("token not an *Token.Token"))
		return
	}
	//如果是admin角色，直接放行
	if tk.Role == user.ROLE_ADMIN {
		return
	}
	if tk.Role != a.role {
		response.Failed(c, exception.NewPermissionDeny("role %d not allow", tk.Role))

		return
	}
}

// 携带参数的gin中间件
func Required(r user.Role) gin.HandlerFunc {
	a := NewTokenAuther()
	a.role = r
	return a.Perm
}
