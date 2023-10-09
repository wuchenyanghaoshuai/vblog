package impl_test

import (
	"context"
	"testing"

	"github.com/wuchenyanghaoshuai/vblog/apps/token"

	"github.com/wuchenyanghaoshuai/vblog/ioc"
	"github.com/wuchenyanghaoshuai/vblog/test"
)

var (
	tokenSvc token.Service
	ctx      = context.Background()
)

func init() {
	test.DevelopmentSetup()
	tokenSvc = ioc.Controller().Get(token.AppName).(token.Service)
}

func TestLoginRequest(t *testing.T) {
	req := token.NewLoginRequest()
	req.Username = "admin2"
	req.Password = "12345"
	tk, err := tokenSvc.Login(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tk)
}

func TestValidateToken(t *testing.T) {
	req := token.NewValidateToken("ckhqphfnl53dbri9vob0")
	tk, err := tokenSvc.ValidateToken(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tk)
}
