package impl

import (
	"context"
	"vblog/apps/blog"
)

func(i *UserServiceImpl) QueryBlog(ctx context.Context, in *blog.QueryBlogRequest) (*blog.BlogSet, error) {
	return nil, nil
}
//文章详情
func(i *UserServiceImpl) DescribeBlog(ctx context.Context, in *blog.DescribeBlogRequest) (*blog.Blog, error) {
	return nil, nil
}
//文章创建
func(i *UserServiceImpl) CreateBlog(ctx context.Context, in *blog.CreateBlogRequest) (*blog.Blog, error) {
	return nil, nil
}
//文章更新
func(i *UserServiceImpl) UpdateBlog(ctx context.Context, in *blog.UpdateBlogRequest) (*blog.Blog, error) {
	return nil, nil
}
//文章删除
func(i *UserServiceImpl) DeleteBlog(ctx context.Context, in *blog.DeleteBlogRequest) (*blog.Blog, error) {
	return nil, nil
}