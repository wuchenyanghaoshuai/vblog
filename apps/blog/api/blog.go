package api

import (
	"vblog/apps/blog"
	"vblog/apps/token"
	"vblog/common"
	"vblog/exception"
	"vblog/middleware"
	"vblog/response"

	"github.com/gin-gonic/gin"
)

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
//postman里直接发起Get请求 
//http://127.0.0.1:8088/vblog/api/v1/blogs?keywords=全站开发
func (h *BlogApiHandler) QueryBlog(c *gin.Context) {
	// 1. 获取用户请求
	req := blog.NewQueryBlogRequest()
	req.PageRequest = common.NewPageRequestFromGinCtx(c)
	req.KeyWords = c.Query("keywords")
	// 2. 业务处理
	set, err := h.svc.QueryBlog(c.Request.Context(), req)
	if err != nil {
		response.Failed(err, c)
		return
	}
	// 3. 返回结果
	response.Success(set, c)
}
func (h *BlogApiHandler) CreateBlog(c *gin.Context) {
	//1. 获取用户请求
	req := blog.NewCreateBlogRequest()
	if err := c.ShouldBindJSON(req); err != nil {
		response.Failed(err, c)
		return
	}
	//补充用户数据
	if  v,ok := c.Get(token.GIN_TOKEN_KEY_NAME);ok{
		req.CreateBy = v.(*token.Token).UserName
	}
	
	//2. 业务处理
	ins, err := h.svc.CreateBlog(c.Request.Context(), req)
	if err != nil {
		response.Failed(err, c)
		return
	}
	//3. 返回结果
	response.Success(ins, c)
}
func (h *BlogApiHandler) DescribeBlog(c *gin.Context) {
	//1. 获取用户请求
	req := blog.NewDescribeBlogRequest(c.Param("id"))
	//2. 业务处理
	ins, err := h.svc.DescribeBlog(c.Request.Context(), req)
	if err != nil {
		response.Failed(err, c)
		return
	}
	//3. 返回结果
	response.Success(ins, c)
}
func (h *BlogApiHandler) PutUpdateBlog(c *gin.Context) {
	//1. 获取用户请求
	req := blog.NewUpdateBlogRequest(c.Param("id"))
	req.UpdateMode = common.UPDATE_MODE_PUT

	//body
	if err := c.Bind(req.CreateBlogRequest); err != nil {
		response.Failed(exception.ErrValidateFailed(err.Error()), c)
		return
	}
	//2. 业务处理

	ins, err := h.svc.UpdateBlog(c, req)
	if err != nil {
		response.Failed(err, c)
		return
	}
	//3. 返回结果
	response.Success(ins, c)
}
func (h *BlogApiHandler) PatchUpdateBlog(c *gin.Context) {
	//1. 获取用户请求
	req := blog.NewUpdateBlogRequest(c.Param("id"))
	req.UpdateMode = common.UPDATE_MODE_PATCH
	//body
	if err := c.Bind(req.CreateBlogRequest); err != nil {
		response.Failed(exception.ErrValidateFailed(err.Error()), c)
		return
	}
	//2. 业务处理
	ins, err := h.svc.UpdateBlog(c, req)
	if err != nil {
		response.Failed(err, c)
		return
	}
	//3. 返回结果
	response.Success(ins, c)
}
func (h *BlogApiHandler) UpdateBlogStatus(c *gin.Context) {
	//1. 获取用户请求
	req := blog.NewUpdateBlogStatusRequest(c.Param("id"))
	//body
	if err := c.Bind(req.ChangeBlogStatusRequest); err != nil {
		response.Failed(exception.ErrValidateFailed(err.Error()), c)
		return
	}
	//2. 业务处理
	ins, err := h.svc.UpdateBlogStatus(c.Request.Context(), req)
	if err != nil {
		response.Failed(err, c)
		return	
	}
	//3. 返回结果
	response.Success(ins, c)
}
func (h *BlogApiHandler) DeleteBlog(c *gin.Context) {
	//1. 获取用户请求
	req := blog.NewDeleteBlogRequest(c.Param("id"))

	//2. 业务处理
	ins,err := h.svc.DeleteBlog(c.Request.Context(), req)
	if err != nil {
		response.Failed(err, c)
		return
	}
	//3. 返回结果
	response.Success(ins, c)
}
