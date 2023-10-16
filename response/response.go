package response

//主要用来重新新定义数据返回格式
// 200 (OK) api 有响应并返回数据 {code:0,data:xxx,msg:xxx}
// http code 沿用 http code规范
// 正常请求的api数据返回200
// 异常情况的api返回 非200
import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wuchenyanghaoshuai/vblog/exception"
)

// 正常情况数据返回
func Success(c *gin.Context, data any) {
	c.JSON(http.StatusOK, data)
}

func Failed(c *gin.Context, err error) {
	defer c.Abort()
	var e *exception.ApiException
	if v, ok := err.(*exception.ApiException); ok {
		e = v
	} else {
		e = exception.New(http.StatusInternalServerError, err.Error())
		e.HttpCode = http.StatusInternalServerError
	}
	c.JSON(e.HttpCode, e)
}
