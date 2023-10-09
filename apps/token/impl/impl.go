package impl

import (
	"context"

	"github.com/wuchenyanghaoshuai/vblog/apps/token"
	"github.com/wuchenyanghaoshuai/vblog/apps/user"
	"github.com/wuchenyanghaoshuai/vblog/conf"
	"github.com/wuchenyanghaoshuai/vblog/exception"
	"github.com/wuchenyanghaoshuai/vblog/ioc"
	"gorm.io/gorm"
)


func init(){
	ioc.Controller().Registry(&TokenServiceImpl{})
}

type TokenServiceImpl struct {
	db *gorm.DB
	//依赖user模块，直接操作user模块的数据库users？
	// 这里需要依赖另一个业务凌虚： 用户管理领域
	user user.Service
}


func(i *TokenServiceImpl) Init()error{
	i.db = conf.C().Mysql.GetConn().Debug()
	i.user =  ioc.Controller().Get(user.AppName).(user.Service)
	return nil
}
func (i *TokenServiceImpl)Name()string{
	return token.AppName
}

func(i *TokenServiceImpl) Login(ctx context.Context,req *token.LoginRequest)(*token.Token,error){
	//1. 查询用户是否存在
	//i.user.DescribeUserRequest(ctx,user.NewDescribeUserRequestByUsername(req.Username))
	//通过user模块，来查询用户是否存在，可以写成上面一行的方式也可以抽成2行
	ureq := user.NewDescribeUserRequestByUsername(req.Username)
	u,err := i.user.DescribeUserRequest(ctx,ureq)
	if err !=nil {
		if exception.IsNotFound(err){
			
			
			return nil,token.AuthFailed
		}
		return nil,err
	}
	
	
	//2. 比对用户传递过来的密码
	err = u.CheckPassword(req.Password)
	if err !=nil{
		return nil,token.AuthFailed
	}

	//3. 颁发token
	tk := token.NewToken()
	tk.UserId = u.Id
	tk.UserName = u.Username
	//4. 颁发完token以后把之前的token标记为失效 （作业） 避免同一个用户多次登陆

	//5. 保存token
	if err := i.db.WithContext(ctx).Create(tk).Error;err !=nil{
		return nil,err
	}
	return tk,nil
}




func(i *TokenServiceImpl) Logout(ctx context.Context,req *token.LogoutRequest)(error){
	return nil
}


func(i *TokenServiceImpl) ValidateToken(ctx context.Context,req *token.ValidateToken)(*token.Token,error){
	//1.查询token （是否是我们这个系统颁发的）
	tk:= token.NewToken()
	if err := i.db.WithContext(ctx).Where("access_token=?",req.AccessToken).First(tk).Error;err !=nil{
		return nil,err
	}
	tk.TableName()
	//2.判断token合法性
	//2.1 判断access_token是不是过期了
	if err := tk.IsExpired();err !=nil{
		return nil,err
	}
	return tk,nil
}