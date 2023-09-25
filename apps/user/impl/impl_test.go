package impl_test

import (
	"context"
	"testing"

	"github.com/wuchenyanghaoshuai/vblog/apps/user"
	"github.com/wuchenyanghaoshuai/vblog/apps/user/impl"
)

var (
	userSvc *impl.UserServiceImpl
	ctx     = context.Background()
)

func init() {
	userSvc = impl.NewUserServiceImpl()
}

func TestCreateUser(t *testing.T) {
	req := user.NewCreateUserRequest()
	req.Username = "admin22"
	req.Password = "12345"
	u, err := userSvc.CreateUser(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(u)
}
func TestDeleteUser(t *testing.T) {
	err := userSvc.DeleteUser(ctx, &user.DeleteUserRequest{Id: 18})
	if err != nil {
		t.Fatal(err)
	}
}

func TestDescribeUserRequestById(t *testing.T) {
	req := user.NewDescribeUserRequestById("15")
	u, err := userSvc.DescribeUserRequest(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(u)
}

func TestDescribeUserRequestByUserName(t *testing.T) {
	req := user.NewDescribeUserRequestByUsername("admin6")
	u, err := userSvc.DescribeUserRequest(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(u)
}
