package ioc

//专门用于注册controller对象
func Controller()*IocContainer{
	return controllerContainer
}

//ioc注册表，全局只有一个

var controllerContainer = &IocContainer{
	store: map[string]IocObject{},
}