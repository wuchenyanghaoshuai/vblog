package impl_test

import (
	"github.com/wuchenyanghaoshuai/vblog/apps/blog"
	"testing"
)

func TestCreateBlog(t *testing.T) {
	in := blog.NewCreateBlogRequest()
	in.Title = "Vblog Web Service Api2"
	in.Content = "Golang"
	in.Tags["分类"] = "Golang"

	ins, err := svc.CreateBlog(ctx, in)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

func TestQueryBlog(t *testing.T) {
	in := blog.NewQueryBlogRequest()
	in.SetStatus(blog.STATUS_PUBLISHED)
	set, err := svc.QueryBlog(ctx, in)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(set)
}

func TestDescribeBlog(t *testing.T) {
	in := blog.NewDescribeBlogRequest("48")
	set, err := svc.DescribeBlog(ctx, in)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(set)
}

func TestUpdateBlogPut(t *testing.T) {
	in := blog.NewPutUpdateBlogRequest("47")
	in.Author = "san"
	set, err := svc.UpdateBlog(ctx, in)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(set)
}

func TestUpdateBlogPatch(t *testing.T) {
	in := blog.NewPatchUpdateBlogRequest("50")
	in.Author = "张鑫2"
	set, err := svc.UpdateBlog(ctx, in)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(set)
}

func TestDeleteBlog(t *testing.T) {
	in := blog.NewDeleteBlogRequest("48")
	err := svc.DeleteBlog(ctx, in)
	if err != nil {
		t.Fatal(err)
	}

}
