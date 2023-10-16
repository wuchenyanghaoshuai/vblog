package token

import (
	"encoding/json"
	"github.com/wuchenyanghaoshuai/vblog/apps/user"
	"time"

	"github.com/rs/xid"
	"github.com/wuchenyanghaoshuai/vblog/exception"
)

func NewToken() *Token {
	return &Token{

		//用它生成一个唯一的uuid的字符串
		AccessToken:           xid.New().String(),
		AccessTokenExpiredAt:  7200,
		RefreshToken:          xid.New().String(),
		RefreshTokenExpiredAt: 3600 * 24 * 7,
		CreatedAt:             time.Now().Unix(),
	}
}

type Token struct {
	//token颁发给谁
	UserId int64 `json:"user_id"`
	//人的名字
	UserName string `json:"username" gorm:"column:username"`
	//颁发给用户的访问令牌(用户需要携带token来访问接口)
	AccessToken string `json:"access_token"`
	//过期时间,单位是秒
	AccessTokenExpiredAt int `json:"access_token_expired_at"`
	//刷新token
	RefreshToken string `json:"refresh_token"`
	//刷新token的过期时间
	RefreshTokenExpiredAt int `json:"refresh_token_expired_at"`
	//创建时间
	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`
	//额外补充信息，忽略gorm对他进行处理
	Role user.Role `gorm:"-"`
}

func (t *Token) TableName() string {
	return "tokens"
}

func (t *Token) IsExpired() error {
	duration := time.Since(t.ExpiredTime())
	expiredSecond := duration.Seconds()
	if expiredSecond > 0 {
		return exception.NewTokenExpired("toekn %s 过期了 %f 秒", t.AccessToken, expiredSecond)
	}
	return nil
}

func (t *Token) ExpiredTime() time.Time {
	return time.Unix(t.CreatedAt, 0).Add(time.Duration(t.AccessTokenExpiredAt) * time.Second)
}

func (u *Token) String() string {
	dj, _ := json.Marshal(u)
	return string(dj)
}
