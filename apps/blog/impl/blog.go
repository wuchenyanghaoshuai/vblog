package impl

import (
	"context"
	"vblog/apps/blog"
	"vblog/common"
	"vblog/exception"

	"dario.cat/mergo"
)

// 文章列表查询
func (i *BlogServiceImpl) QueryBlog(ctx context.Context, in *blog.QueryBlogRequest) (*blog.BlogSet, error) {
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

// 文章详情
func (i *BlogServiceImpl) DescribeBlog(ctx context.Context, in *blog.DescribeBlogRequest) (*blog.Blog, error) {
	ins := blog.NewBlog()
	err := i.db.WithContext(ctx).Where("id = ?", in.BlogId).Table("blogs").First(ins).Error
	if err != nil {
		return nil, err
	}
	return ins, nil
}

// 文章创建
func (i *BlogServiceImpl) CreateBlog(ctx context.Context, in *blog.CreateBlogRequest) (*blog.Blog, error) {
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

// 文章更新 比较大，支持全量更新以及部分更新
func (i *BlogServiceImpl) UpdateBlog(ctx context.Context, in *blog.UpdateBlogRequest) (*blog.Blog, error) {
	//先把要更新的对象查询出来
	ins, err := i.DescribeBlog(ctx, blog.NewDescribeBlogRequest(in.BlogId))
	if err != nil {
		return nil, err
	}
	//使用switch来判断更新模式
	switch in.UpdateMode {
	case common.UPDATE_MODE_PUT:
		//全量更新 ,你传递什么我就保存什么
		ins.CreateBlogRequest = in.CreateBlogRequest
	case common.UPDATE_MODE_PATCH:

		//部分更新	， 只更新有变化的字段
		err := mergo.MergeWithOverwrite(ins.CreateBlogRequest, in.CreateBlogRequest)
		if err != nil {
			return nil, err
		}
	}
	//更新字段校验
	if err := ins.CreateBlogRequest.Validate(); err != nil {
		return nil, exception.ErrValidateFailed(err.Error())
	}
	//更新数据
	err = i.db.WithContext(ctx).Table("blogs").Save(ins).Error
	if err != nil {
		return nil, err
	}
	return ins, nil
}

// 文章删除
func (i *BlogServiceImpl) DeleteBlog(ctx context.Context, in *blog.DeleteBlogRequest) (*blog.Blog, error) {
	ins, err := i.DescribeBlog(ctx, blog.NewDescribeBlogRequest(in.BlogId))
	if err != nil {
		return nil, err
	}

	//删除数据
	err = i.db.WithContext(ctx).Table("blogs").Where("id = ?", in.BlogId).Delete(ins).Error
	if err != nil {
		return nil, err
	}

	return ins, nil
}

// 文章发布
func (i *BlogServiceImpl) UpdateBlogStatus(ctx context.Context, in *blog.UpdateBlogStatusRequest) (*blog.Blog, error) {

	ins, err := i.DescribeBlog(ctx, blog.NewDescribeBlogRequest(in.BlogId))
	if err != nil {
		return nil, err
	}
	//更新状态
	ins.ChangeBlogStatusRequest = in.ChangeBlogStatusRequest
	ins.SetStatus(in.Status) // 设置状态
	err = i.db.WithContext(ctx).Table("blogs").Where("id = ?", in.BlogId).
		Updates(map[string]interface{}{"status": in.Status}).Error
	if err != nil {
		return nil, err
	}
	//返回更新后的实例
	return ins, nil
}
