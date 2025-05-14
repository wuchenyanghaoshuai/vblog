# IOC (依赖管理) 对象的依赖管理

```go
	userServiceImpl :=user.NewUserServiceImpl()
	tokenServiceImpl := token.NewTokenServiceImpl(userServiceImpl)
	tokenApiHandler := api.NewTokenApiHandler(tokenServiceImpl)
```
在程序启动的时候(main.go) 收到依赖传递，main组装流程复杂



#IOC
让软件工程规模化，很好的解决了工程依赖关系

#IOC Container

系统内所有业务对象(BO)的一个托儿所，这样业务对象的依赖才能自己去IOC Container要