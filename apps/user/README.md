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