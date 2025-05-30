package impl_test

import (
	"testing"
	"vblog/apps/blog"
	"vblog/common"
)


func TestCreateBlog(t *testing.T) {
	req := blog.NewCreateBlogRequest()
	req.Author = "Chenyang Wu"
	req.Title = "Go全站开发5"
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
	req := blog.NewDescribeBlogRequest("5")
	ins,err := serviceImpl.DescribeBlog(ctx, req)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(ins)
}
//UPDATE `blogs` SET `created_at`=1748426490,`updated_at`=1748505588,`title`='更新后的文章标题',`author`='Chenyang Wu',`content`='MD内容填充',`summary`='Go全站开发的内容简介',`tags`='{"test":"test"}' WHERE `id` = 5
// UPDATE `blogs` SET `created_at`=1748426490,`updated_at`=1748506182,`title`='更新后的文章标题2',`author`='Chenyang Wu',`content`='MD内容填充',`summary`='Go全站开发的内容简介',`tags`='{"test":"test"}' WHERE `id` = 5
//只更新了部分字段，也就是title字段
func TestPatchUpdateBlog(t *testing.T) {
	req := blog.NewUpdateBlogRequest("5")
	req.UpdateMode = common.UPDATE_MODE_PATCH
	req.Title="更新后的文章标题2"
	ins,err := serviceImpl.UpdateBlog(ctx, req)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(ins)
}
//UPDATE `blogs` SET `created_at`=1748426490,`updated_at`=1748506676,`title`='put更新后的文章标题',`author`='Chenyang Wu',`content`='put',`tags`='{}' WHERE `id` = 5
func TestPutUpdateBlog(t *testing.T) {
	req := blog.NewUpdateBlogRequest("5")
	req.UpdateMode = common.UPDATE_MODE_PUT
	req.Title = "put更新后的文章标题"
	req.Author = "Chenyang Wu"
	req.Content = "put"
	ins,err := serviceImpl.UpdateBlog(ctx, req)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(ins)
}

//DELETE FROM `blogs` WHERE `blogs`.`id` = 5
//DELETE FROM `blogs` WHERE id = '4' AND `blogs`.`id` = 4
func TestDeleteBlog(t *testing.T) {
	req := blog.NewDeleteBlogRequest("4")
	ins,err := serviceImpl.DeleteBlog(ctx, req)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(ins)
}

func TestUpdateBlogStatus(t *testing.T) {
	req := blog.NewUpdateBlogStatusRequest("3")
	req.SetStatus(blog.Status_Published) // 设置状态为已发布
	
	ins,err := serviceImpl.UpdateBlogStatus(ctx, req)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(ins)
}