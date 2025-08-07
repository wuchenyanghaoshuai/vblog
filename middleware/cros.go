package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Cros(c *gin.Context) {
	if c.Request.Method == http.MethodOptions {
		c.Header("Access-Control-Allow-Origin","*")
		c.Header("Access-Control-Allow-Methods","GET,POST,PUT,DELETE,OPTIONS")
		c.Header("Access-Control-Allow-Headers","*")
		c.Header("Access-Control-Allow-Credentials","true")
		c.Header("Access-Control-Max-Age","86400")
		c.Header("Content-Type","application/json")
		c.Header("Content-Length","0")
		c.Header("Content-Disposition","inline")
		c.Header("Content-Transfer-Encoding","binary")
		c.Header("Content-Encoding","gzip")
		c.Writer.WriteHeader(http.StatusNoContent)
		return
	}
	c.Next()
}