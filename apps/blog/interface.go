package blog

import (
	"context"
	"vblog/common"
)

// AppName 应用名称
const AppName = "blogs"

type Service interface {
	//文章列表查询
	QueryBlog(context.Context, *QueryBlogRequest) (*BlogSet, error)
	//文章详情
	DescribeBlog(context.Context, *DescribeBlogRequest) (*Blog, error)
	//文章创建
	CreateBlog(context.Context,*CreateBlogRequest) (*Blog, error)
	//文章更新
	UpdateBlog(context.Context,*UpdateBlogRequest) (*Blog, error)
	//文章发布
	UpdateBlogStatus(context.Context, *UpdateBlogStatusRequest) (*Blog, error)
	//文章删除
	DeleteBlog(context.Context,*DeleteBlogRequest) (*Blog,error)
}
func NewUpdateBlogStatusRequest(BlogId string) *UpdateBlogStatusRequest {
	return &UpdateBlogStatusRequest{
		BlogId:                 BlogId,
		ChangeBlogStatusRequest: &ChangeBlogStatusRequest{},
	}
}
type UpdateBlogStatusRequest struct {
	//文章ID
	BlogId string `json:"blog_id" validate:"required"`
	//文章状态
	*ChangeBlogStatusRequest
}
func NewQueryBlogRequest() *QueryBlogRequest {
	return &QueryBlogRequest{
		PageRequest: common.NewPageRequest(),
		KeyWords:    "",
	}
}

type QueryBlogRequest struct {
	*common.PageRequest

	//关键字参数，根据文章名称包含某个字
	KeyWords string
}
func NewDescribeBlogRequest(id string) *DescribeBlogRequest {
	return &DescribeBlogRequest{
		BlogId: id,
	}
}
type DescribeBlogRequest struct {
	//文章ID
	BlogId string `json:"blog_id"`
}

func NewUpdateBlogRequest(BlogId string) *UpdateBlogRequest {
	return &UpdateBlogRequest{
		BlogId:      BlogId,
		UpdateMode:  common.UPDATE_MODE_PUT, //默认全量更新
		CreateBlogRequest: NewCreateBlogRequest(),
	}
}
func (req *UpdateBlogRequest) Validate() error {
	return common.Validate(req)
}

type UpdateBlogRequest struct{
	//文章ID
	BlogId string `json:"blog_id" validate:"required"`
	//更新模型，全量更新还是部分更新
	UpdateMode common.UPDATE_MODE  `json:"update_mode"`
	//需要更新的数据
	*CreateBlogRequest `validate:"required"` //使用嵌套结构体来复用CreateBlogRequest的字段
}


func NewDeleteBlogRequest(id string) *DeleteBlogRequest {
	return &DeleteBlogRequest{
		BlogId: id,
	}
}

type DeleteBlogRequest struct {
	//文章ID
	BlogId string `json:"blog_id"`
}