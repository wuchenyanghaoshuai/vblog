package token

import (
	"encoding/json"
	"time"
	"vblog/apps/user"

	"github.com/rs/xid"
)

func NewToken(u *user.User) *Token {
	return &Token{
		UserId: u.Id,
		UserName: u.Username,
		//生成随机字符串
		AccessToken: xid.New().String(),
		AccessTokenExpiredAt: 3600,
		RefreshToken: xid.New().String(),
		RefreshTokenExpiredAt: 3600*4,
		CreatedAt: time.Now().Unix(),
		Role: u.Role,
	}
}
func DefaultToken() *Token {
	return &Token{}
}
type Token struct {
	UserId int `json:"user_id" gorm:"column:user_id"`
	UserName string `json:"user_name"  gorm:"column:username"`
	AccessToken string `json:"access_token" gorm:"column:access_token"`
	AccessTokenExpiredAt int `json:"access_token_expired_at" gorm:"column:access_token_expired_at"`
	RefreshToken string `json:"refresh_token" gorm:"column:refresh_token"`
	RefreshTokenExpiredAt int `json:"refresh_token_expired_at" gorm:"column:refresh_token_expired_at"`
	CreatedAt int64 `json:"created_at" gorm:"column:created_at"`
	UpdatedAt int `json:"updated_at" gorm:"column:updated_at"`
	Role user.Role `json:"role" gorm:"-"`
}
func (t *Token) IssueTime() time.Time {
	return time.Unix(t.CreatedAt,0)
}

func (t *Token)AccessTokenDuration() time.Duration {
	return time.Duration(t.AccessTokenExpiredAt) * time.Second
}

func (t *Token)RefreshTokenDuration() time.Duration {
	return time.Duration(t.RefreshTokenExpiredAt) * time.Second
}

func (t *Token)AccessTokenIsExpired() error {
	//过期时间 = 颁发时间+过期时长
	expiredTime := t.IssueTime().Add(t.AccessTokenDuration())
	if time.Since(expiredTime).Seconds() > 0 {
		return ErrAccessTokenExpired
	}
	return nil
}
func (t *Token)RefreshokenIsExpired() error {
	expiredTime := t.IssueTime().Add(t.RefreshTokenDuration())
	if time.Since(expiredTime).Seconds() > 0 {
		return ErrRefreshTokenExpired
	}
	return nil
}

func (t *Token)String() string {
	dj,_ :=json.MarshalIndent(t, "", "    ")
	return string(dj)
}