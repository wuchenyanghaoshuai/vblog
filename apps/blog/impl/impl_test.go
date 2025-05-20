package impl_test

import (
	"context"
	"vblog/apps/blog"
	"vblog/ioc"

	//倒入被测试的对象，全部倒入
	_ "vblog/apps"
)

var (
	serviceImpl blog.Service
	ctx         = context.Background()
)

func init() {
	serviceImpl = ioc.Controller.Get(blog.AppName).(blog.Service)
}