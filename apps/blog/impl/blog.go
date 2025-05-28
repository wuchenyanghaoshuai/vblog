package impl

import (
	"context"
	"vblog/apps/blog"
	"vblog/exception"
)

//文章列表查询
func(i *BlogServiceImpl) QueryBlog(ctx context.Context, in *blog.QueryBlogRequest) (*blog.BlogSet, error) {
	set := blog.NewBlogSet()
	//1. 有默认值， 不需要用户传递参数
	// 2.构造查询条件  关键字匹配
	query := i.db.WithContext(ctx).Table("blogs")
	if in.KeyWords != "" {
		query = query.Where("title LIKE ? OR content LIKE ?", "%"+in.KeyWords+"%", "%"+in.KeyWords+"%")
	}
	//3.查询总数
	err := query.Count(&set.Total).Error
	if err != nil {
		return nil, err
	}
	//查询
	err = query.Limit(in.PageSize).Offset(in.Offset()).Find(&set.Items).Error
	if err != nil {
		return nil, err
	}
	return set, nil
}
//文章详情
func(i *BlogServiceImpl) DescribeBlog(ctx context.Context, in *blog.DescribeBlogRequest) (*blog.Blog, error) {
	ins := blog.NewBlog()

	err := i.db.WithContext(ctx).Table("blogs").First(ins, in.BlogId).Error
	if err != nil {
		return nil, err
	}
	return ins, nil
}
//文章创建
func(i *BlogServiceImpl) CreateBlog(ctx context.Context, in *blog.CreateBlogRequest) (*blog.Blog, error) {
	//1. 验证请求参数
	if err := in.Validate(); err != nil {
		return nil, exception.ErrValidateFailed(err.Error())
	}
	//2. 构造查询实例
	ins := blog.NewBlog()
	ins.CreateBlogRequest = in

	//3. 入库返回
	//INSERT INTO `blogs` (`created_at`,`updated_at`,`title`,`author`,`content`,`summary`,`create_by`,`tags`,`published_at`,`status`) VALUES (1748405000,1748405000,'Go全站开发','Chenyang Wu','MD内容填充','Go全站开发的内容简介','','{"test":"test"}',0,0)
	err := i.db.WithContext(ctx).Create(ins).Error
	if err != nil {
		return nil, err
	}

	return ins, nil
}
//文章更新
func(i *BlogServiceImpl) UpdateBlog(ctx context.Context, in *blog.UpdateBlogRequest) (*blog.Blog, error) {
	return nil, nil
}
//文章删除
func(i *BlogServiceImpl) DeleteBlog(ctx context.Context, in *blog.DeleteBlogRequest) (*blog.Blog, error) {
	return nil, nil
}

//文章发布
func(i *BlogServiceImpl) UpdateBlogStatus(ctx context.Context, in *blog.ChangeBlogStatusRequest) (*blog.Blog, error) {

	return nil, nil
}