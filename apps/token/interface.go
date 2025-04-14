package token

import "context"


type Service interface {
	//令牌颁发
	IssueToken(context.Context,*IssueToken) (*Token, error)
	//令牌撤销
	RevolkToken(context.Context,*RevolkToken) (*Token, error)
	//令牌校验
	ValidateToken(context.Context,*ValidateToken) (*Token, error)
}

type IssueToken struct {
	Username string
	Password string
	Ismember bool
}

type RevolkToken struct {
	//你需要撤销的令牌
	// accesstoken跟refreshToken 构成了一对username/password
	AccessToken string
	//你需要知道正确的刷新令牌
	RefreshToken string
}
type ValidateToken struct {
	//你需要撤销的令牌
	// accesstoken跟refreshToken 构成了一对username/password
	AccessToken string
}