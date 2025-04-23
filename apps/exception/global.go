package exception

import "fmt"

func ErrServerInternal(format string,a ...any) *ApiException {
	return &ApiException{
		Code: 50000,
		Message: fmt.Sprintf(format,a...),
	}
}