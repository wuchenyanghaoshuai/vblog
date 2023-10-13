package impl_test

import (
	"context"
	"github.com/wuchenyanghaoshuai/vblog/apps/blog"
	"github.com/wuchenyanghaoshuai/vblog/ioc"
	"github.com/wuchenyanghaoshuai/vblog/test"
)

var (
	svc blog.Service
	ctx = context.Background()
)

func init() {
	test.DevelopmentSetup()

	// 依赖另一个实现类
	svc = ioc.Controller().Get(blog.AppName).(blog.Service)
}
