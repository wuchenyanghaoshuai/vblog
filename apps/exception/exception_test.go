package exception_test

import (
	"testing"
	"vblog/apps/exception"
)




func ApiException()error {
	return exception.NewApiException(50001, "token失效")
}


func TestXxx(t *testing.T) {
	e :=ApiException()
	t.Log(e)
	
	// 类型断言：将 error 转换为 *ApiException
    apiErr, ok := e.(*exception.ApiException)
    if !ok {
        t.Fatal("Expected *ApiException")
    }
    t.Logf("Code: %d, Message: %s", apiErr.Code, apiErr.Message)
}