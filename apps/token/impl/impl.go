package impl

import (
	"context"
	"fmt"
	"vblog/apps/conf"
	"vblog/apps/exception"
	"vblog/apps/ioc"
	"vblog/apps/token"
	"vblog/apps/user"

	"gorm.io/gorm"
)

//
func init(){
	//注册到ioc容器中
	ioc.Controller.Registry(token.AppName,&TokenServiceImpl{})
}




// func NewTokenServiceImpl(userserviceimpl user.Service) *TokenServiceImpl {
// 	return &TokenServiceImpl{
// 		db: conf.C().MySQL.GetDB(),
// 		user: userserviceimpl,
// 	}
// }

func(i *TokenServiceImpl) Init()error{
	i.db = conf.C().MySQL.GetDB()
	//获取对象
	//断言对象
	i.user = ioc.Controller.Get(user.AppName).(user.Service)
	return nil
}

type TokenServiceImpl struct {
	db *gorm.DB
	user user.Service
}

func(i *TokenServiceImpl) IssueToken(ctx context.Context,in *token.IssueTokenRequest) (*token.Token, error){
	//查询用户对象
	// queryuser 是调用的user.service的接口吗
	queryUser := user.NewQueryUserRequest()
	queryUser.Username = in.Username
	us,err := i.user.QueryUser(ctx,queryUser)
	if err != nil {
		return nil,err
	}

	if len(us.Items) == 0 {
		return nil,token.ErrAuthFailed
	}

	//比对用户传递的密码和数据库的密码是否一致
	if err := us.Items[0].CheckPassword(in.Password);err != nil {
		return nil,token.ErrAuthFailed
	}
	//颁发令牌
	tk := token.NewToken(us.Items[0])
	//存储令牌到数据库
	if err := i.db.WithContext(ctx).Create(tk).Error;err!= nil {
		return nil,exception.ErrServerInternal("保存报错, %s",err)
	}

	//返回令牌
	return tk,nil
}
//令牌撤销
func (i *TokenServiceImpl) RevolkToken(ctx context.Context,in *token.RevolkTokenRequest) (*token.Token,error){
	tk := token.DefaultToken()
	//查询出token
	err := i.db.WithContext(ctx).Where("access_token=?",in.AccessToken).First(tk).Error
	if err!= nil {
		return nil,exception.ErrServerInternal("查询token报错, %s",err)
	}

	if tk.RefreshToken != in.RefreshToken {
		return nil,fmt.Errorf("RefreshToken不匹配")
	}

	//直接删除数据库中的令牌
	err = i.db.WithContext(ctx).Where("access_token=?",in.AccessToken).Delete(token.Token{}).Error
	if  err !=nil{
		return nil,err
	}
	return tk,nil
	
}
//令牌校验
func (i *TokenServiceImpl) ValidateToken(ctx context.Context,in *token.ValidateTokenRequest) (*token.Token, error){
	// 查询出token
	tk := token.DefaultToken()
	err := i.db.WithContext(ctx).Where("access_token=?",in.AccessToken).First(tk).Error
	if err != nil {
		return nil,exception.ErrServerInternal("查询token报错, %s",err)
	}
	//判断token是否过期
	if err := tk.RefreshokenIsExpired();err!= nil {
		return nil,err
	}
	if err := tk.AccessTokenIsExpired();err!= nil {
		return nil,err
	}
	//返回token
	return tk,nil
}