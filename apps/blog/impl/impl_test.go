package impl_test

import (
	"context"
	"vblog/apps/blog"
	"vblog/ioc"
	_ "vblog/apps"
)

var (
	serviceImpl blog.Service
	ctx         = context.Background()
)

func init() {
	serviceImpl = ioc.Controller.Get(blog.AppName).(blog.Service)
}