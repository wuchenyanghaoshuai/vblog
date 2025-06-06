package token

import (
	"net/http"
	"vblog/exception"
)

var (
	 ErrAuthFailed = exception.NewApiException(50001, "token失效").WithHttpCode(http.StatusUnauthorized)
	 ErrAccessTokenExpired = exception.NewApiException(50002, "AccessToken已过期")
	 ErrRefreshTokenExpired = exception.NewApiException(50003, "RefreshToken已过期")
	 ErrPermissionDenied = exception.NewApiException(50004, "权限不足,访问未授权,请联系管理员").WithHttpCode(http.StatusForbidden)

)