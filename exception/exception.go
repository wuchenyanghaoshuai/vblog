package exception

import (
	"fmt"
	"net/http"
)

func New(code int, format string, a ...any) *ApiException {
	HttpCode := http.StatusInternalServerError
	if code/100 < 6 && code/100 > 0 {
		HttpCode = code
	}
	return &ApiException{
		BizCode:  code,
		Message:  fmt.Sprintf(format, a...),
		HttpCode: HttpCode,
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
