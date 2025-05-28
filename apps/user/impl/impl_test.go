package impl_test

import (
	"context"
	"crypto/md5"
	"fmt"
	"testing"
	"vblog/apps/user"
	"vblog/ioc"
	"vblog/test"

	"golang.org/x/crypto/bcrypt"
)

var (
	serviceImpl user.Service
	ctx         = context.Background()
)

//	CreateUser(context.Context,*CreateUserRequest)(*User,error)
//	QueryUser(context.Context,*QueryUserRequest)(*[]UserSet,error)

func init() {
	//初始化单测环境
	test.DevelopmentSetup()
	//serviceImpl = impl.NewUserServiceImpl()
	serviceImpl = ioc.Controller.Get(user.AppName).(user.Service)
}

func TestCreateUser(t *testing.T) {

	req := user.NewCreateUserRequest()
	req.Username = "adminnnn"
	req.Password = "123456"
	//Label map[string]string
	req.Label = map[string]string{
		"角色": "管理员",
	}
	req.Role=1
	
	ins, err := serviceImpl.CreateUser(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

func TestQueryUser(t *testing.T) {
	req := user.NewQueryUserRequest()
	req.Username = "admin"
	ins, err := serviceImpl.QueryUser(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}
func TestMd5(t *testing.T){
	h := md5.New()
	h.Write([]byte("123456"))
	fmt.Printf("%x\n", h.Sum(nil))
	bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
}

//$2a$14$JCx/rGNV0PjXdNLEL6okhuFGTpUX6iZOc0G9v0fFTfREBac3m.YBW
//$2a$14$07OwLqgxuSFTN24hYHDqsOqy155WrhWdkDNfzYmB3bksl7TKeDm1y
func TestPasswordHash(t *testing.T) {
	password := "secret"
	hash, _ := HashPassword(password) // ignore error for the sake of simplicity

	fmt.Println("Password:", password)
	fmt.Println("Hash:    ", hash)

	match := CheckPasswordHash(password, hash)
	fmt.Println("Match:   ", match)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
func TestUserCheckPassword(t *testing.T) {
	req := user.NewCreateUserRequest()
	req.Username = "admin"
	req.Password = "123456"
	u := user.NewUser(req)

	u.HashPassword()
	t.Log(u.CheckPassword("123456"))
}