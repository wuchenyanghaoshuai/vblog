package impl

import (
	"context"
	"dario.cat/mergo"
	"fmt"
	"github.com/wuchenyanghaoshuai/vblog/apps/blog"
	"time"
)

func (i *blogServiceImpl) CreateBlog(ctx context.Context, req *blog.CreateBlogRequest) (*blog.Blog, error) {
	ins := blog.NewBlog(req)
	if err := i.db.WithContext(ctx).Create(ins).Error; err != nil {
		return nil, err
	}
	return ins, nil
}

func (i *blogServiceImpl) UpdateBlogStatus(ctx context.Context, in *blog.UpdateBlogStatusRequest) (*blog.Blog, error) {
	return nil, nil
}

func (i *blogServiceImpl) UpdateBlog(ctx context.Context, in *blog.UpdateBlogRequest) (*blog.Blog, error) {
	ins, err := i.DescribeBlog(ctx, blog.NewDescribeBlogRequest(in.BlogId))
	if err != nil {
		return nil, err
	}
	switch in.UpdateMode {
	case blog.UPDATE_MODE_PUT:
		//全量更新
		ins.CreateBlogRequest = in.CreateBlogRequest
	case blog.UPDATE_MODE_PATCH:
		//增量更新 就是对比传递过来的数据跟数据库里的对比，如果有就更新，没有就按照原来的
		err := mergo.Merge(ins.CreateBlogRequest, in.CreateBlogRequest, mergo.WithOverride)
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("unknown update mode:%d", in.UpdateMode)
	}
	//更新数据
	ins.UpdatedAt = time.Now().Unix()
	fmt.Println("打印打印", ins)
	err = i.db.WithContext(ctx).Where("id=?", in.BlogId).Updates(ins).Error
	if err != nil {
		return nil, err
	}
	return ins, nil
}

func (i *blogServiceImpl) DeleteBlog(ctx context.Context, in *blog.DeleteBlogRequest) error {

	return i.db.WithContext(ctx).Table("blogs").Where("id=?", in.BlogId).Delete(&blog.Blog{}).Error
}

func (i *blogServiceImpl) QueryBlog(ctx context.Context, in *blog.QueryBlogRequest) (*blog.BlogSet, error) {

	query := i.db.WithContext(ctx).Model(&blog.Blog{})
	//提前准备好set对象
	set := blog.NewBlogSet()
	//组装查询条件
	if in.Status != nil {
		query = query.Where("status = ?", *in.Status)
	}

	//1.查询总数量  SELECT * FROM `blogs` WHERE status = '1' LIMIT 10
	err := query.Count(&set.Total).Error
	if err != nil {
		return nil, err
	}

	//2.查询一页的数据
	err = query.Offset(in.Offset()).Limit(int(in.PageSize)).Find(&set.Items).Error
	if err != nil {
		return nil, err
	}
	return set, nil
}

func (i *blogServiceImpl) AuditBlog(ctx context.Context, in *blog.AuditBlogRequest) (*blog.Blog, error) {
	ins, err := i.DescribeBlog(ctx, blog.NewDescribeBlogRequest(in.BlogId))
	if err != nil {
		return nil, err
	}
	ins.IsAuditPass = in.IsAuditPass

	ins.AuditAt = time.Now().Unix()
	err = i.db.WithContext(ctx).Where("id=?", in.BlogId).Updates(ins).Error
	if err != nil {
		return nil, err
	}
	return ins, err
}

func (i *blogServiceImpl) DescribeBlog(ctx context.Context, in *blog.DescribeBlogRequest) (*blog.Blog, error) {
	query := i.db.WithContext(ctx).Model(&blog.Blog{})
	ins := blog.NewBlog(blog.NewCreateBlogRequest())
	err := query.Where("id = ?", in.BlogId).First(ins).Error
	if err != nil {
		return nil, err
	}
	return ins, nil
}
