package blog_test

import (
	"testing"
	"time"
	"vblog/apps/blog"
) 

func TestBlog(t *testing.T) {
	ins := blog.NewBlog()
	t.Log(ins)
}

func TestCreateBlogRequest(t *testing.T) {
	ins := blog.NewBlog()
	ins.CreateBlogRequest.Title = "test"
	ins.CreateBlogRequest.Content = "tis is a blog's content"
	ins.CreateBlogRequest.Summary = "tis is a blog's summary"
	ins.CreateBlogRequest.CreateBy = "admin"
	ins.CreateBlogRequest.Author = "chenyang wu"
	ins.CreateBlogRequest.Tags = map[string]string{"历史": "文艺复兴"}
	t.Log(ins.CreateBlogRequest)
}

func TestChangeBlogStatusRequest(t *testing.T) {
	ins := blog.NewBlog()
	ins.PublishAt = time.Now().Unix()
	ins.Status = "0"
	t.Log(ins.ChangeBlogStatusRequest)
}