# Gin Web 中间件
## 中间件(存在于 请求跟响应之间)

##  Auth认证
```
auth 认证解决的是识别用户的问题，比如登录用户为a或者b
```
### 定义中间件
```go
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
```

###  加载中间件

```go
注意有先后顺序
func (h *BlogApiHandler) Registry(appRouter gin.IRouter) {
	//不需要鉴权的接口
	appRouter.GET("/", h.QueryBlog)
	appRouter.GET("/:id", h.DescribeBlog)
	//修改变更需要认证
	appRouter.Use(middleware.Auth)
	appRouter.POST("/", h.CreateBlog)
	appRouter.PUT("/:id", h.PutUpdateBlog)
	appRouter.PATCH("/:id", h.PatchUpdateBlog)
	appRouter.POST("/:id/status", h.UpdateBlogStatus)
	appRouter.DELETE("/:id", h.DeleteBlog)
}
```

###  获取中间件当中注入的上下文
```go
	//获取中间件当中注入的上下文
	if  v,ok := c.Get(token.GIN_TOKEN_KEY_NAME);ok{
		req.CreateBy = v.(*token.Token).UserName
	}
```

## 鉴权
```
用户登录以后，你能干什么，你已经登录，但并不是所有的功能你都可以用
```

