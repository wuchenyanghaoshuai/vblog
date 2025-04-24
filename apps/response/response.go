package response

import (
	"net/http"
	"vblog/apps/exception"

	"github.com/gin-gonic/gin"
)

//统一对外返回
//这个也可以做脱敏处理
//就不用在每一个c.json里写了
func Success(data any,c *gin.Context){
	c.JSON(http.StatusOK,data)
}



func Failed(err error,c *gin.Context){
	//返回一个非200的状态码 返回内容Apiexception
	httpCode := http.StatusInternalServerError
	
	if apiErr,ok := err.(*exception.ApiException);ok{
		if apiErr.Code != 0 {
		httpCode = apiErr.HttpCode
		}
	}else{
	 err = exception.ErrServerInternal(err.Error())
	}
	c.JSON(httpCode, err)
	c.Abort()
}