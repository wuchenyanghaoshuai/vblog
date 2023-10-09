package apps

//业务控制器（负责倒入所有的业务实现）注册到ioc连的controller区域

import(
	// 倒入包的先后顺序就是对象注册的先后顺序
	_ "github.com/wuchenyanghaoshuai/vblog/apps/user/impl"
	_ "github.com/wuchenyanghaoshuai/vblog/apps/token/impl"
	_ "github.com/wuchenyanghaoshuai/vblog/apps/token/api"
)