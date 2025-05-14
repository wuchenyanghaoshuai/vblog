package ioc

var Controller Container = &MapContainer{
	storage: make(map[string]Object),
}


