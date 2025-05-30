package api

import (
	"vblog/apps/blog"
	"vblog/conf"
	"vblog/ioc"
)



func init() {
	ioc.Api.Registry(blog.AppName, &BlogApiHandler{})
}

func (h *BlogApiHandler) Init() error {
	 h.svc = ioc.Controller.Get(blog.AppName).(blog.Service)
	 subRouter := conf.C().Application.GinRootRouter().Group("blogs")
	 h.Registry(subRouter)
	return nil
}

type BlogApiHandler struct {
	svc blog.Service
}