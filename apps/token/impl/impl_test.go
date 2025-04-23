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
	req:=token.NewIssueTokenRequest("admin","123456")
	tk,err := serviceImpl.IssueToken(ctx,req)
	if err != nil  {
		log.Fatal(err)
	}
	t.Log(tk)
}


func TestRevolkToken(t *testing.T) {
	req := token.NewRevolkTokenRequest("d04a86emrvq71f4q9ar0","d04a86emrvq71f4q9arg")
	tk,err := serviceImpl.RevolkToken(ctx,req)
	if err!= nil {
		log.Fatal(err)
	}
	t.Log(tk)
}

func TestValidateToken(t *testing.T) {
	req := token.NewValidateTokenRequest("d04a86emrvq71f4q9ar0")
	tk,err :=serviceImpl.ValidateToken(ctx,req)
	if err!= nil {
		log.Fatal(err)
	}
	t.Log(tk)
}