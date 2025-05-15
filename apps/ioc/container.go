package ioc

//Controller 是一个container，使用mapcontainer实现
var Controller Container = &MapContainer{
	name:    "Controller",
	storage: make(map[string]Object),
}


//APi所有的对外接口对象都存放在这里

var Api Container = &MapContainer{
	name:    "Api",
	storage: make(map[string]Object),
}