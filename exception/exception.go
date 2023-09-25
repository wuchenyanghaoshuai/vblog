package exception

import "fmt"


func New(code int,format string,a ...any)*ApiException{
	return &ApiException{
		Code: code,
		Message: fmt.Sprintf(format, a...),
	}
}

type ApiException struct{
	Code int `json:"code"`
	Message string `json:"message"`
}

func (e *ApiException)Error()string{
	return e.Message
}