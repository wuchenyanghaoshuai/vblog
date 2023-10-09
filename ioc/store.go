package ioc

import "github.com/gin-gonic/gin"

// 定义一个接口,就是定义业务逻辑
// 定义一个对象的注册表

type IocContainer struct {
	//采用map来保存对象的注册关系
	store map[string]IocObject
}

func (c *IocContainer) Init() error {
	for _, obj := range c.store {
		if err := obj.Init(); err != nil {
			return err
		}
	}
	return nil
}

func (c *IocContainer) Registry(obj IocObject) {
	c.store[obj.Name()] = obj
}
func (c *IocContainer) Get(name string) any {
	return c.store[name]
}

type GinApiHandler interface {
	Registry(r gin.IRouter)
}

func (c *IocContainer) RouteRegistry(r gin.IRouter) {
	for _, obj := range c.store {
		if api, ok := obj.(GinApiHandler); ok {
			api.Registry(r)
		}
	}
}
