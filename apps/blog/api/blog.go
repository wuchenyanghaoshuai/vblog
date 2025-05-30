package api

import (
	"vblog/apps/blog"
	"vblog/common"
	"vblog/response"

	"github.com/gin-gonic/gin"
)


func(h *BlogApiHandler)Registry(appRouter gin.IRouter){
	appRouter.POST("/",h.CreateBlog)
	appRouter.GET("/",h.QueryBlog)
	appRouter.GET("/:id", h.DescribeBlog)
	appRouter.PUT("/:id", h.PutUpdateBlog)
	appRouter.PATCH("/:id", h.PatchUpdateBlog)
	appRouter.POST("/:id/status", h.UpdateBlogStatus)
	appRouter.DELETE("/:id", h.DeleteBlog)
}

func (h *BlogApiHandler) QueryBlog(c *gin.Context) {
	// 1. 获取用户请求
	req := blog.NewQueryBlogRequest()	
	req.PageRequest = common.NewPageRequestFromGinCtx(c)
	req.KeyWords = c.Query("keywords")
	// 2. 业务处理
	set,err := h.svc.QueryBlog(c.Request.Context(), req)
	if err != nil {
		response.Failed(err, c)
		return
	}
	// 3. 返回结果
	response.Success(set, c)
}
func (h *BlogApiHandler) CreateBlog(c *gin.Context) {
}
func (h *BlogApiHandler) DescribeBlog(c *gin.Context) {
}
func (h *BlogApiHandler) PutUpdateBlog(c *gin.Context) {
}
func (h *BlogApiHandler) PatchUpdateBlog(c *gin.Context) {
}
func (h *BlogApiHandler) UpdateBlogStatus(c *gin.Context) {
}
func (h *BlogApiHandler) DeleteBlog(c *gin.Context) {
}
