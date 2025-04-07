package user

import (
	"encoding/json"
	"vblog/apps/common"
)

//里面存放的是po，也就是需要持久化的东西

func NewCreateUserRequest()*CreateUserRequest {
	return &CreateUserRequest{
		Role: Role_VISITOR,
		Label: map[string]string{},
	}
}
//需要用户输入的信息
type CreateUserRequest struct {
	UserName string `json:"username" gorm:"column:username" validate:"required"`
	Password string `json:"password" gorm:"column:password" validate:"required"`
	Role  Role `json:"role" gorm:"column:role"`
	//用户标签,本身map不支持直接存到数据库中，需要进行序列化
	// 可以使用gorm的serializer标签，将map序列化为json字符串
	// https://gorm.io/docs/serializer.html
	Label map[string]string `json:"label" gorm:"column:label;serializer:json"`
}


func NewUser(req *CreateUserRequest)*User {
	return &User{
		Meta: common.NewMeta(),
		CreateUserRequest: req,
	}
}
func (req *User) String() string {
	dj,_ := json.MarshalIndent(req, "", "    ")
	return string(dj)
}

//通用参数
type Meta struct {
	//用户id
	UserId int `json:"user_id" gorm:"column:user_id"`
	//创建时间
	CreatedAt int64 `json:"created_at" gorm:"column:created_at"`
	//更新时间
	UpdatedAt int64 `json:"updated_at" gorm:"column:updated_at"`
}

//用户创建成功以后返回一个user对象
type User struct {
	*common.Meta
	//用户信息
	*CreateUserRequest
}

func NewUserSet()*UserSet {
	return &UserSet{
		Items: []*User{},
	}
}

type UserSet struct {
	Total int `json:"total"`
	Items []*User `json:"items"`
}