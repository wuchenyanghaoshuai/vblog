package token

import "context"


type Service interface {
	//令牌颁发
	//调用user模块校验用户密码
	IssueToken(context.Context,*IssueTokenRequest) (*Token, error)
	//令牌撤销
	//删除令牌
	RevolkToken(context.Context,*RevolkTokenRequest) (*Token, error)
	//令牌校验
	//检查令牌是否是我们自己颁发的
	//并且没有过期
	ValidateToken(context.Context,*ValidateTokenRequest) (*Token, error)
}

func NewIssueTokenRequest(username,password string) *IssueTokenRequest {
	return &IssueTokenRequest{
		Username: username,
		Password: password,
		Ismember: false,
	}
}
type IssueTokenRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Ismember bool   `json:"is_member"`
}
func NewRevolkTokenRequest(accessToken,refreshToken string) *RevolkTokenRequest {
	return &RevolkTokenRequest{
		AccessToken: accessToken,
		RefreshToken: refreshToken,
	}
}
type RevolkTokenRequest struct {
	//你需要撤销的令牌
	// accesstoken跟refreshToken 构成了一对username/password
	AccessToken string
	//你需要知道正确的刷新令牌
	RefreshToken string
}
type ValidateTokenRequest struct {
	//你需要撤销的令牌
	// accesstoken跟refreshToken 构成了一对username/password
	AccessToken string
}

func NewValidateTokenRequest(accessToken string) *ValidateTokenRequest {
	return &ValidateTokenRequest{
		AccessToken: accessToken,
	}
}