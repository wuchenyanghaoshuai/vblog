package token

import "vblog/exception"

var (
	 ErrAuthFailed = exception.NewApiException(50001, "token失效")
	 ErrAccessTokenExpired = exception.NewApiException(50002, "AccessToken已过期")
	 ErrRefreshTokenExpired = exception.NewApiException(50003, "RefreshToken已过期")
)