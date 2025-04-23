package impl_test

import (
	"context"
	"log"
	"testing"
	"vblog/apps/token"
	"vblog/apps/token/impl"
	user "vblog/apps/user/impl"
)


var (
  serviceImpl token.Service
  ctx = context.Background()	
)

func init() {
	serviceImpl = impl.NewTokenServiceImpl(user.NewUserServiceImpl())
}

func TestIssueToken(t *testing.T) {
	req:=token.NewIssueTokenRequest("admin","1234567")
	tk,err := serviceImpl.IssueToken(ctx,req)
	if err != nil  {
		log.Fatal(err)
	}
	t.Log(tk)
}


func TestRevolkToken(t *testing.T) {
	serviceImpl.RevolkToken(ctx,nil)
}

func TestValidateToken(t *testing.T) {
	serviceImpl.ValidateToken(ctx,nil)
}