package user

import (
	"context"


	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// 定义user包的能力,就是定义接口
// 站在使用方的角度来定义，
// 接口定义好了不要试图随意修改接口，要保证接口的兼容性
type Service interface {
	CreateUser(context.Context, *CreateUserRequest) (*User, error)
	DeleteUser(context.Context, *DeleteUserRequest) error
	DescribeUserRequest(context.Context,*DescribeUserRequest)(*User,error)
}

// 创建用户请求
type CreateUserRequest struct {
	Username string `json:"username" gorm:"column:username"`
	Password string `json:"password"`
	Role     Role   `json:"role"`
	//比如就存储在数据库里面 存储为json，需要orm来帮我们完成json的序列化和存储
	Label map[string]string `json:"label" gorm:"serializer:json"`
	isHashed bool
}

func NewCreateUserRequest() *CreateUserRequest{
	return &CreateUserRequest{Role: ROLE_MEMBER,Label: map[string]string{}}
}

func (req *CreateUserRequest) PasswordHash(){
	if req.isHashed{
		return
	}
	b, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	req.Password = string(b)
	req.isHashed = true
}


func(req *CreateUserRequest) SetIsHashed(){
	req.isHashed = true
}

func(req *CreateUserRequest) Validate()error{
	if req.Username==""||req.Password==""{
		
		return fmt.Errorf("用户名密码不能为空")
	}
	return nil
}

// 删除用户请求
type DeleteUserRequest struct {
	Id int64 `json:"id"`
}


func NewDescribeUserRequestById(id string) *DescribeUserRequest {
	return &DescribeUserRequest{
		DescribeValue: id,
	}
}

func NewDescribeUserRequestByUsername(username string) *DescribeUserRequest {
	return &DescribeUserRequest{
		DescribeBy:    DESCRIBE_BY_USERNAME,
		DescribeValue: username,
	}
}

//查询用户 要支持 id 和username来查询
type DescribeUserRequest struct{
	DescribeBy    DescribeBy `json:"describe_by"`
	DescribeValue string     `json:"describe_value"`
}

