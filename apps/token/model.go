package token

import "encoding/json"

type Token struct {
	UserId string `json:"user_id" gorm:"column:user_id"`
	UserName string `json:"user_name"  gorm:"column:username"`
	AccessToken string `json:"access_token" gorm:"column:access_token"`
	AccessTokenExpiredAt int `json:"access_token_expired_at" gorm:"column:access_token_expired_at"`
	RefreshToken string `json:"refresh_token" gorm:"column:refresh_token"`
	RefreshTokenExpiredAt int `json:"refresh_token_expired_at" gorm:"column:refresh_token_expired_at"`
	CreatedAt int `json:"created_at" gorm:"column:created_at"`
	UpdatedAt int `json:"updated_at" gorm:"column:updated_at"`
	Role string `json:"role" gorm:"-"`
}

func (t *Token)String() string {
	dj,_ :=json.MarshalIndent(t, "", "    ")
	return string(dj)
}