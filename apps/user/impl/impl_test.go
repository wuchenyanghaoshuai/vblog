package impl_test

import (
	"context"
	"testing"

	"github.com/wuchenyanghaoshuai/vblog/apps/user"

	"github.com/wuchenyanghaoshuai/vblog/ioc"
	"github.com/wuchenyanghaoshuai/vblog/test"
)

var (
	userSvc user.Service
	ctx     = context.Background()
)

func init() {
	test.DevelopmentSetup()
	//取出对象
	//断言为接口来使用（只使用对象接口提供出来的能力）
	userSvc = ioc.Controller().Get(user.AppName).(user.Service)
}

func TestCreateUser(t *testing.T) {
	req := user.NewCreateUserRequest()
	req.Username = "admin222"
	req.Password = "12345"
	u, err := userSvc.CreateUser(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(u)
}
func TestDeleteUser(t *testing.T) {
	err := userSvc.DeleteUser(ctx, &user.DeleteUserRequest{Id: 26})
	if err != nil {
		t.Fatal(err)
	}
}

func TestDescribeUserRequestById(t *testing.T) {
	req := user.NewDescribeUserRequestById("22")
	u, err := userSvc.DescribeUserRequest(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(u)
}

func TestDescribeUserRequestByUserName(t *testing.T) {
	req := user.NewDescribeUserRequestByUsername("admin1")
	u, err := userSvc.DescribeUserRequest(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(u)
}
