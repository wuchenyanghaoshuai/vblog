package user

import (
	"encoding/json"

	"github.com/wuchenyanghaoshuai/vblog/common"
	"golang.org/x/crypto/bcrypt"
)

// 用于存放  存入数据库的对象
type User struct {
	//这里添加其他的字段，跟用户传递过来的做一个组合，然后统一存入数据库
	//例如用户传递过来了账号密码，我这边加一些 创建时间 之类的字段一起存入数据库
	//通用信息复用
	*common.Meta
	//集合用户传递过来的数据(用户名和密码)
	*CreateUserRequest
}

func (u *User)String()string{
	dj,_ :=json.Marshal(u)
	return string(dj)
}

//判断该用户的密码是否正确
func(u *User) CheckPassword(password string)error{
	return bcrypt.CompareHashAndPassword([]byte(u.Password),[]byte(password))
}


func NewUser(req *CreateUserRequest)*User{
	req.PasswordHash()
	return &User{
		Meta: common.NewMeta(),
		CreateUserRequest: req,
	}
}

//声明你这个队形存储在users表里面
//orm 负责调用tbalename（）来动态获取你这个对象要存储的表的名称
func (u *User) TableName()string{
	return "users"
}