package token_test

import (
	"vblog/apps/token"
	"testing"
)


func TestTokenString(t *testing.T) {
	token := &token.Token{
		UserId: "1",
		UserName: "1",
		AccessToken: "1",
		AccessTokenExpiredAt: 1,
		RefreshToken: "1",
		RefreshTokenExpiredAt: 1,
		CreatedAt: 1,
		UpdatedAt: 1,
		Role: "1",
	}
//	tokenString := token.String()
	
	t.Log(token.String())
}