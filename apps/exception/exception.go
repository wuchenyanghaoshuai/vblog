package exception

//用于描述业务异常
type ApiException struct {
	//业务异常编码，例如 50001  表示token失效
	Code int `json:"code"`
	Message  string `json:"message"`
	//
	HttpCode int `json:"-"`
}

//如何实现error接口
func (e *ApiException) Error() string {
	return e.Message
}
func NewApiException(code int, message string) *ApiException {
	return &ApiException{
		Code: code,
		Message: message,
	}
}