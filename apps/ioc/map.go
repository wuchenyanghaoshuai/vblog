package ioc

type MapContainer struct {
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
	for _ ,v := range c.storage{
		if err := v.Init();err != nil {
			return err
		}
	}
	return nil
}