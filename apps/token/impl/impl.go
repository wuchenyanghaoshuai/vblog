package impl

import (
	"context"
	"vblog/apps/conf"
	"vblog/apps/token"
	"vblog/apps/user"

	"gorm.io/gorm"
)

func NewTokenServiceImpl(userserviceimpl user.Service) *TokenServiceImpl {
	return &TokenServiceImpl{
		db: conf.C().MySQL.GetDB(),
		user: userserviceimpl,
	}
}

type TokenServiceImpl struct {
	db *gorm.DB
	user user.Service
}

func(i *TokenServiceImpl) IssueToken(ctx context.Context,in *token.IssueToken) (*token.Token, error){
	//查询用户对象
	// queryuser 是调用的user.service的接口吗
	queryUser := user.NewQueryUserRequest()
	queryUser.Username = in.Username
	us,err := i.user.QueryUser(ctx,queryUser)
	if err != nil {
		return nil,err
	}

	//比对用户传递的密码和数据库的密码是否一致
	us.Items[0].CheckPassword(in.Password)
	//颁发令牌

	//存储令牌到数据库

	//返回令牌
	return nil,nil
}
//令牌撤销
func (i *TokenServiceImpl) RevolkToken(context.Context,*token.RevolkToken) (*token.Token, error){
	//直接删除数据库中的令牌
	return nil,nil
}
//令牌校验
func (i *TokenServiceImpl) ValidateToken(context.Context,*token.ValidateToken) (*token.Token, error){
	// 查询出token
	//判断token是否过期
	//返回token
	return nil,nil
}