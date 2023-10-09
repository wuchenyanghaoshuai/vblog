package ioc
//专门用于注册controller对象
func ApiHandler()*IocContainer{
	return apiHandlerContainer
}

//ioc注册表，全局只有一个

var apiHandlerContainer = &IocContainer{
	store: map[string]IocObject{},
}