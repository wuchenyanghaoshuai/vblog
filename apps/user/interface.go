package user

import (
	"context"
	"vblog/common"
)

const (
	//业务包的名称,用于托管业务包的业务对象
	AppName = "user"
)
type Service interface{
	//创建用户
	//用户取消了请求怎么办
	// 如果做trace，reaceid怎么传递
	// 多个接口，需要做事物(Session)
	CreateUser(context.Context,*CreateUserRequest)(*User,error)
	//用户查询
	QueryUser(context.Context,*QueryUserRequest)(*UserSet,error)

}

func NewQueryUserRequest()*QueryUserRequest{
	return &QueryUserRequest{
		PageRequest: common.NewPageRequest(),
	}
}

//不需要持久化写入到数据库，所以就不放到model里面
type QueryUserRequest struct{
	Username string
	*common.PageRequest
}