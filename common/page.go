package common

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func NewPageRequest()*PageRequest {
	return &PageRequest{
		PageSize: 10,
		PageNum: 1,
	}
}
func NewPageRequestFromGinCtx(c *gin.Context) *PageRequest {
	p := NewPageRequest()
	pnStr := c.Query("page_number")
	psStr := c.Query("page_size")
	if pnStr != "" {
		pn,_:=strconv.ParseInt(pnStr, 10, 64)
		if pn !=0 {
			p.PageNum = int(pn)
		}
	if psStr != "" {
		ps, _ := strconv.ParseInt(psStr, 10, 64)
		if ps != 0 {
			p.PageSize = int(ps)
		}
		}
	}
	return p
}

type PageRequest struct {
	//分页大小
	PageSize int `json:"page_size"`
	//第几页
	PageNum int `json:"page_num"`
}

func (req *PageRequest) Offset() int {
	return (req.PageNum - 1) * req.PageSize
}