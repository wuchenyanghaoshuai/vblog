package impl

import (
	"vblog/apps/blog"
	"vblog/conf"
	"vblog/ioc"

	"gorm.io/gorm"
)


func init(){
	ioc.Controller.Registry(blog.AppName, &UserServiceImpl{})
}

type UserServiceImpl struct {
	db *gorm.DB

}

func (i *UserServiceImpl) Init() error {

	i.db = conf.C().MySQL.GetDB()
	return nil
}