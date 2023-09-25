package impl_test

import (
	"context"
	"testing"

	"github.com/wuchenyanghaoshuai/vblog/apps/token"
	"github.com/wuchenyanghaoshuai/vblog/apps/token/impl"
	userImpl "github.com/wuchenyanghaoshuai/vblog/apps/user/impl"
)

var (
	tokenSvc *impl.TokenServiceImpl
	ctx      = context.Background()
)

func init() {
	tokenSvc = impl.NewTokenServiceImpl(userImpl.NewUserServiceImpl())
}

func TestLoginRequest(t *testing.T) {
	req := token.NewLoginRequest()
	req.Username = "admin22"
	req.Password = "12345"
	tk, err := tokenSvc.Login(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tk)
}

func TestValidateToken(t *testing.T) {
	req := token.NewValidateToken("ck6g7sbjtoj4ak9orlvg")
	tk, err := tokenSvc.ValidateToken(ctx,req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tk)
}