# 用户模块

## 接口定义(业务定义)
1. 定义User对象结构
```go
//用户创建成功以后返回一个user对象
type User struct {
	*common.Meta
	//用户信息
	*CreateUserRequest
}
```
2. 定义接口

```go
type Service interface{
	//创建用户
	//用户取消了请求怎么办
	// 如果做trace，reaceid怎么传递
	// 多个接口，需要做事物(Session)
	CreateUser(context.Context,*CreateUserRequest)(*User,error)
	//用户查询
	QueryUser(context.Context,*QueryUserRequest)(*[]UserSet,error)
}

```
3. 接口的具体实现



### 密码存储问题