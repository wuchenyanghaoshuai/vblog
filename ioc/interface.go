package ioc

//定义注册进来的对象的约束条件


type IocObject interface{
	Init() error
	Name()string
}