package ioc

import "fmt"



type MapContainer struct {
	name string
	storage map[string]Object
}

//注册对象
func(c *MapContainer)Registry(name string, obj Object){
	c.storage[name] = obj
}
//获取对象
func(c *MapContainer)Get(name string)any{
	return c.storage[name]
}
func (c *MapContainer)Init()error{
	for k ,v := range c.storage{
		if err := v.Init();err != nil {
			return fmt.Errorf("init object %s failed, err: %v", k, err)
		}
		fmt.Printf("[%s] %s init success \n",c.name, k)
	}
	return nil
}


