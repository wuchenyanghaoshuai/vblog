package exception

import (
	"fmt"
	"net/http"
)

func New(code int, format string, a ...any) *ApiException {
	return &ApiException{
		BizCode:  code,
		Message:  fmt.Sprintf(format, a...),
		HttpCode: http.StatusOK,
	}
}

type ApiException struct {
	//业务异常
	BizCode  int    `json:"code"`
	Message  string `json:"message"`
	Data     any    `json:"data"`
	HttpCode int    `json:"http_code"`
}

func (e *ApiException) Error() string {
	return e.Message
}
