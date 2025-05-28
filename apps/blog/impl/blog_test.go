package impl_test

import (
	"testing"
	"vblog/apps/blog"
)


func TestCreateBlog(t *testing.T) {
	req := blog.NewCreateBlogRequest()
	req.Author = "Chenyang Wu"
	req.Title = "Go全站开发3"
	req.Content = "MD内容填充"
	req.Summary = "Go全站开发的内容简介"
	//Tags map[string]string `json:"tags" gorm:"column:tags;serializer:json"`
	req.Tags["test"] = "test"
	ins,err := serviceImpl.CreateBlog(ctx, req)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(ins)
}
func TestQueryBlog(t *testing.T) {
	req := blog.NewQueryBlogRequest()
	req.KeyWords = "Go全站开发"
	ins,err := serviceImpl.QueryBlog(ctx, req)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(ins)
}

func TestDescribeBlog(t *testing.T) {
	req := blog.NewDescribeBlogRequest("1")
	ins,err := serviceImpl.DescribeBlog(ctx, req)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(ins)
}