package token_test

import (
	"testing"
	"time"
	"vblog/apps/token"
	"vblog/apps/user"
)


func TestTokenString(t *testing.T) {
	token := &token.Token{
		UserId: 1,
		UserName: "1",
		AccessToken: "1",
		AccessTokenExpiredAt: 1,
		RefreshToken: "1",
		RefreshTokenExpiredAt: 1,
		CreatedAt: 1,
		UpdatedAt: 1,
		Role: user.Role_Admin,
	}
//	tokenString := token.String()
	
	t.Log(token.String())
}


func TestTokenExired(t *testing.T) {
	now := time.Now().Unix()
	tk := token.Token{
		UserId: 1,
		Role: user.Role_Admin,
		AccessTokenExpiredAt: 1,
		CreatedAt: now,
	}
	t.Log(tk.AccessTokenIsExpired())
}