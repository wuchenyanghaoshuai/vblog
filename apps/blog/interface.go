package blog

import (
	"context"
	"encoding/json"
	"strconv"
)

const (
	AppName = "blogs"
)

type Service interface {
	//创建博客
	CreateBlog(context.Context, *CreateBlogRequest) (*Blog, error)
	//发布博客
	UpdateBlogStatus(context.Context, *UpdateBlogStatusRequest) (*Blog, error)
	//查询博客列表,列表查询 没有必要查询文章的具体内容
	QueryBlog(context.Context, *QueryBlogRequest) (*BlogSet, error)
	//详情页，尽量多的把关联数据查出来
	DescribeBlog(context.Context, *DescribeBlogRequest) (*Blog, error)
	//更新博客
	UpdateBlog(context.Context, *UpdateBlogRequest) (*Blog, error)
	//删除博客
	DeleteBlog(context.Context, *DeleteBlogRequest) error
	//文章审核，审核通过的菜可以被看到
	AuditBlog(context.Context, *AuditBlogRequest) (*Blog, error)
}

func NewAuditBlogRequest(id string) *AuditBlogRequest {
	return &AuditBlogRequest{
		BlogId: id,
	}
}

type AuditBlogRequest struct {
	BlogId string `json:"blog_id"`
	//是否审核成功
	IsAuditPass bool `json:"is_audit_pass"`
}

func NewDescribeBlogRequest(id string) *DescribeBlogRequest {
	return &DescribeBlogRequest{
		BlogId: id,
	}
}

type DescribeBlogRequest struct {
	BlogId string `json:"blog_id"`
}

func NewBlogSet() *BlogSet {
	return &BlogSet{
		Items: []*Blog{},
	}
}

type BlogSet struct {
	//博客总数
	Total int64 `json:"total"`
	//返回一页的数据
	Items []*Blog `json:"items"`
}

func (b *BlogSet) String() string {
	dj, _ := json.Marshal(b)
	return string(dj)
}

func (s *BlogSet) Add(items ...*Blog) {
	s.Items = append(s.Items, items...)
}

func (r *QueryBlogRequest) Offset() int {
	return int(r.PageSize * (r.PageNumber - 1))
}

func (r *QueryBlogRequest) PasePageSize(ps string) {
	psInt, _ := strconv.ParseInt(ps, 10, 64)
	if psInt != 0 {
		r.PageSize = int(psInt)
	}
}
func (r *QueryBlogRequest) PasePageNumber(pn string) {
	psInt, _ := strconv.ParseInt(pn, 10, 64)
	if psInt != 0 {
		r.PageNumber = int(psInt)
	}
}

func (r *QueryBlogRequest) SetStatus(s Status) {
	r.Status = &s
}
func NewQueryBlogRequest() *QueryBlogRequest {
	return &QueryBlogRequest{
		PageSize:   10,
		PageNumber: 1,
	}
}

type QueryBlogRequest struct {
	//页的大小，一页有多少个内容
	PageSize int `json:"page_size"`
	//当前在第几页
	PageNumber int `json:"page_number"`
	//0 表示草稿状态，要查询所有的博客，nil 没有这个过滤条件，0 --》draft   ；1---》published
	Status *Status `json:"status"`
}

type UpdateBlogStatusRequest struct {
	//如果定义一篇文章，使用对象Id，具体的某一篇文章
	BlogId int64 `json:"blog_id"`
	//修改状态，发布或者草稿
	Status Status `json:"status"`
}

func NewPutUpdateBlogRequest(id string) *UpdateBlogRequest {
	return &UpdateBlogRequest{
		BlogId:            id,
		UpdateMode:        UPDATE_MODE_PUT,
		CreateBlogRequest: NewCreateBlogRequest(),
	}
}
func NewPatchUpdateBlogRequest(id string) *UpdateBlogRequest {
	return &UpdateBlogRequest{
		BlogId:            id,
		UpdateMode:        UPDATE_MODE_PATCH,
		CreateBlogRequest: NewCreateBlogRequest(),
	}
}

// 区分全量更新和部分更新
type UpdateBlogRequest struct {
	BlogId string `json:"blog_id"`
	//更新模式
	UpdateMode UpdateMode `json:"update_mode"`
	*CreateBlogRequest
}

func NewDeleteBlogRequest(id string) *DeleteBlogRequest {
	return &DeleteBlogRequest{
		BlogId: id,
	}
}

type DeleteBlogRequest struct {
	BlogId string `json:"blog_id"`
}
