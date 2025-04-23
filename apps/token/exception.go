package token

import "vblog/apps/exception"

var (
	 ErrAuthFailed = exception.NewApiException(50001, "token失效")
)