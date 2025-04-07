package blog

import (
	"encoding/json"
	"vblog/apps/common"
)

// 为什么这么写
//因为直接使用blog的话在初始化的时候会提供具体的值，如果使用newblog，直接返回了一个*Blog，而不是具体的值
// 因为没有具体的默认值，所以在手动调用或者测试的时候，直接赋值就可以了，最大化的保证了灵活性
//构造函数，对 对象进行初始化，统一对对象的初始化者进行管理
// 保证构造的对象是可用的，不容易出现nil的情况
func NewBlog()*Blog {
	return &Blog{
		Meta: common.NewMeta(),
		CreateBlogRequest: &CreateBlogRequest{
			Tags: map[string]string{},
		},
		ChangeBlogStatusRequest: &ChangeBlogStatusRequest{},
	}
}

type Blog struct {
	*common.Meta
	*CreateBlogRequest
	*ChangeBlogStatusRequest
}
func (req *Blog) String() string {
	dj,_ := json.MarshalIndent(req, "", "    ")
	return string(dj)
}



type CreateBlogRequest struct {
	//标题
	Title string `json:"title" gorm:"column:title" validate:"required"`
	//作者
	Author string `json:"author" gorm:"column:author" validate:"required"`
	//内容
	Content string `json:"content" gorm:"column:content" validate:"required"`
	//摘要
	Summary string `json:"summary" gorm:"column:summary"`
	//创建者
	CreateBy string `json:"create_by" gorm:"column:create_by"`
	Tags map[string]string `json:"tags" gorm:"column:tags;serializer:json"`
}
func (req *CreateBlogRequest) String() string {
	dj,_ := json.MarshalIndent(req, "", "    ")
	return string(dj)
}


type ChangeBlogStatusRequest struct {
	//发布时间
	PublishAt int64 `json:"publish_at" gorm:"column:publish_at"`
	//状态  草稿/已发布
	Status string `json:"status" gorm:"column:status"`
}
func (req *ChangeBlogStatusRequest) String() string {
	dj,_ := json.MarshalIndent(req, "", "    ")
	return string(dj)
}