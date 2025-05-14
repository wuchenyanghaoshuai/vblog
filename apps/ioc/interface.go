package ioc

type Container interface {
	Registry(name string, obj Object)
	Get(name string) any
	Init() error
}


type Object interface {
	Init() error
}