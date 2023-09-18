package context

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

type UserName struct{}
type UserRole struct{}

func Curl(ctx context.Context, url string) error {
	client := http.DefaultClient
	//resp, err := client.Get(url)
	//if err != nil {
	//	return err
	//}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req = req.WithContext(ctx)
	resp, err := client.Do(req)
	//携带上下文

	body, err := io.ReadAll(resp.Body)
	fmt.Printf("%s \n", body)
	return nil
}

// 用户认证Login()
// 转账payment()
func Payment(ctx context.Context, tk string) error {
	ctx1 := context.WithValue(ctx, UserName{}, "bobbbbbb")
	ctx2 := context.WithValue(ctx1, UserRole{}, "userrolekey--system-admin")
	ctx3 := context.WithValue(ctx2, "token", tk)
	DoPayment(ctx3)

	return nil
}

func DoPayment(ctx context.Context) {
	fmt.Println(ctx.Value(UserName{}))
	fmt.Println(ctx.Value(UserRole{}))
	fmt.Println(ctx.Value("token"))

}

//
