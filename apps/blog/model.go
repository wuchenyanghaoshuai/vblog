package blog

import (
	"encoding/json"
	"time"
)

func NewCreateBlogRequest() *CreateBlogRequest {
	return &CreateBlogRequest{
		Tags:    make(map[string]string),
		Title:   "",
		Author:  "",
		Content: "",
		Summary: "",
	}
}

// 用户创建文章
type CreateBlogRequest struct {
	//文章标题
	Title string `json:"title"`
	//文章作者
	Author string `json:"author"`
	//用户登陆后 我们通过token知道是那个人
	CreatedBy string `json:"created_by"`
	//文章内容
	Content string `json:"content"`
	//文章概要
	Summary string `json:"summary"`
	//文章标签,基于标签做分类
	Tags map[string]string `json:"tags" gorm:"serializer:json"`
}

func NewBlog(req *CreateBlogRequest) *Blog {
	return &Blog{
		CreatedAt:         time.Now().Unix(),
		Status:            STATUS_DRAFT,
		CreateBlogRequest: req,
	}
}

type Blog struct {
	//文章id 文章的唯一标识符 给程序使用
	Id int64 `json:"id"`
	//创建时间
	CreatedAt int64 `json:"created_at"`
	//更新时间
	UpdatedAt int64 `json:"updated_at"`
	//审核时间
	AuditAt int64 `json:"audit_at"`
	//发布时间
	PublishedAt int64 `json:"published_at"`
	//文章的状态
	Status Status `json:"status"`
	//是否审核成功
	IsAuditPass bool `json:"is_audit_pass"`
	//用户创建博客参数
	*CreateBlogRequest
}

func (b *Blog) TableName() string {
	return "blogs"
}

func (b *Blog) String() string {
	dj, _ := json.Marshal(b)
	return string(dj)
}
