package user

import (
	"encoding/json"
	"fmt"

	"vblog/apps/common"

	"golang.org/x/crypto/bcrypt"
)

//里面存放的是po，也就是需要持久化的东西

func NewCreateUserRequest() *CreateUserRequest {
	
	return &CreateUserRequest{
		Role:  Role_VISITOR,
		Label: map[string]string{},
	}
}

//需要用户输入的信息
type CreateUserRequest struct {
	Username string `json:"username" gorm:"column:username" validate:"required"`
	Password string `json:"password" gorm:"column:password" validate:"required"`
	Role  Role `json:"role" gorm:"column:role"`
	//用户标签,本身map不支持直接存到数据库中，需要进行序列化
	// 可以使用gorm的serializer标签，将map序列化为json字符串
	// https://gorm.io/docs/serializer.html
	Label map[string]string `json:"label" gorm:"column:label;serializer:json"`
}



//validator 校验器
func (req *CreateUserRequest) Validate() error  {
	if req.Username == "" {
		return fmt.Errorf("用户名不能为空")
	}
	return nil
}

func (req *CreateUserRequest) HashPassword()  error{
	cryptoPass,err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	req.Password = string(cryptoPass)
	return nil
}

func (req *CreateUserRequest) CheckPassword(password string)  error{
	return  bcrypt.CompareHashAndPassword([]byte(req.Password), []byte(password))

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
func(req *User) TableName() string {
	return "users"
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
	Total int64 `json:"total"`
	Items []*User `json:"items"`
}