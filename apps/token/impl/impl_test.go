package impl_test

import (
	"context"
	"log"
	"testing"
	"vblog/apps/token"
	"vblog/test"

	"vblog/ioc"
)


var (


  serviceImpl token.Service
  ctx = context.Background()	
)

func init() {
	//初始化单测环境
	test.DevelopmentSetup()
	//serviceImpl = impl.NewTokenServiceImpl(user.NewUserServiceImpl())
	//去ioc中获取 被测试的业务对象
	serviceImpl = ioc.Controller.Get(token.AppName).(token.Service)
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