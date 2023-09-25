package token

import "context"


type Service interface{
	//登陆接口
	Login(context.Context,*LoginRequest)(*Token,error)
	//退出接口
	Logout(context.Context,*LogoutRequest)error
	//校验token（给内部中间层使用） 身份校验层,
	//校验完返回token，通过token获取用户信息
	ValidateToken(context.Context,*ValidateToken)(*Token,error)
}


type LoginRequest struct{
	Username string `json:"username"`
	Password string `json:"password"`
}
func NewLoginRequest()*LoginRequest{
	return &LoginRequest{}
}
//如果token泄露，不知道refreshtoken也是没办法退出的
type LogoutRequest struct{
	AccessToken string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type ValidateToken struct{
	AccessToken string `json:"access_token"`
}
func NewValidateToken(at string)*ValidateToken{
	return &ValidateToken{
		AccessToken: at,
	}
}