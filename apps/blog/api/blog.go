package api

import (
	"github.com/gin-gonic/gin"
	"github.com/wuchenyanghaoshuai/vblog/apps/blog"
	"github.com/wuchenyanghaoshuai/vblog/apps/token"
	"github.com/wuchenyanghaoshuai/vblog/apps/user"
	"github.com/wuchenyanghaoshuai/vblog/middlewares"
	"github.com/wuchenyanghaoshuai/vblog/response"
)

func (h *apiHandler) Registry(r gin.IRouter) {
	v1 := r.Group("v1").Group("blogs")
	v1.GET("/", h.QueryBlog)
	v1.GET("/:id", h.DescribeBlog)

	//后台管理接口，需要鉴权
	v1.Use(middlewares.NewTokenAuther().Auth)
	v1.POST("/", middlewares.Required(user.ROLE_AUTHOR), h.CreateBlog)
	v1.PUT("/:id", middlewares.Required(user.ROLE_AUTHOR), h.UpdateBlog)
	v1.PATCH("/:id", middlewares.Required(user.ROLE_AUTHOR), h.PatchBlog)
	v1.DELETE("/:id", middlewares.Required(user.ROLE_AUTHOR), h.DeleteBlog)
	v1.POST("/:id/audit", middlewares.Required(user.ROLE_AUDITOR), h.AuditBlog)
}

func (h *apiHandler) CreateBlog(c *gin.Context) {
	//从gin请求上下文中 c.keys 获取认证过后的鉴权结果
	tkObj := c.Keys[token.TOKEN_GIN_KEY_NAME]
	tk := tkObj.(*token.Token)
	in := blog.NewCreateBlogRequest()
	err := c.BindJSON(in)
	if err != nil {
		response.Failed(c, err)
		return
	}
	in.CreatedBy = tk.UserName
	ins, err := h.svc.CreateBlog(c.Request.Context(), in)
	if err != nil {
		response.Failed(c, err)
		return
	}
	response.Success(c, ins)
}
func (h *apiHandler) QueryBlog(c *gin.Context) {

	in := blog.NewQueryBlogRequest()
	in.PasePageSize(c.Query("page_size"))
	in.PasePageNumber(c.Query("page_number"))
	switch c.Query("status") {
	case "draft":
		in.SetStatus(blog.STATUS_DRAFT)
	case "published":
		in.SetStatus(blog.STATUS_PUBLISHED)
	}
	ins, err := h.svc.QueryBlog(c.Request.Context(), in)
	if err != nil {
		response.Failed(c, err)
		return
	}
	response.Success(c, ins)
}

func (h *apiHandler) UpdateBlog(c *gin.Context) {
	in := blog.NewPutUpdateBlogRequest(c.Param("id"))
	err := c.BindJSON(in.CreateBlogRequest)
	if err != nil {
		response.Failed(c, err)
		return
	}
	ins, err := h.svc.UpdateBlog(c.Request.Context(), in)
	if err != nil {
		response.Failed(c, err)
		return
	}
	response.Success(c, ins)
}
func (h *apiHandler) PatchBlog(c *gin.Context) {
	in := blog.NewPatchUpdateBlogRequest(c.Param("id"))
	err := c.BindJSON(in.CreateBlogRequest)
	if err != nil {
		response.Failed(c, err)
		return
	}
	ins, err := h.svc.UpdateBlog(c.Request.Context(), in)
	if err != nil {
		response.Failed(c, err)
		return
	}
	response.Success(c, ins)
}

func (h *apiHandler) DescribeBlog(c *gin.Context) {
	in := blog.NewDescribeBlogRequest(c.Param("id"))
	ins, err := h.svc.DescribeBlog(c.Request.Context(), in)
	if err != nil {
		response.Failed(c, err)
		return
	}
	response.Success(c, ins)
}

func (h *apiHandler) DeleteBlog(c *gin.Context) {
	//判断用户是否登陆
	//判断当前用户是谁，有没有权限删除
	in := blog.NewDeleteBlogRequest(c.Param("id"))
	err := h.svc.DeleteBlog(c.Request.Context(), in)
	if err != nil {
		response.Failed(c, err)
		return
	}
	response.Success(c, "ok")
}

func (h *apiHandler) AuditBlog(c *gin.Context) {
	//判断用户是否登陆
	//判断当前用户是谁，有没有权限删除
	in := blog.NewAuditBlogRequest(c.Param("id"))
	err := c.BindJSON(in)
	if err != nil {
		response.Failed(c, err)
		return
	}
	ins, err := h.svc.AuditBlog(c.Request.Context(), in)
	if err != nil {
		response.Failed(c, err)
		return
	}
	response.Success(c, ins)
}
