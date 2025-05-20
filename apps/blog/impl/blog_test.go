package impl_test

import (
	"testing"
	"vblog/apps/blog"
)


func TestCreateBlog(t *testing.T) {
	req := blog.NewCreateBlogRequest()
	req.Author = "test"
	req.Title = "test"
	req.Content = "test"
	//Tags map[string]string `json:"tags" gorm:"column:tags;serializer:json"`
	req.Tags["test"] = "test"
	
	ins,err := serviceImpl.CreateBlog(ctx, req)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(ins)
}