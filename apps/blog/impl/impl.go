package impl

import (
	"github.com/wuchenyanghaoshuai/vblog/apps/blog"
	"github.com/wuchenyanghaoshuai/vblog/conf"
	"github.com/wuchenyanghaoshuai/vblog/ioc"
	"gorm.io/gorm"
)

func init() {
	ioc.Controller().Registry(&blogServiceImpl{})
}

type blogServiceImpl struct {
	db *gorm.DB
}

func (i *blogServiceImpl) Init() error {
	i.db = conf.C().Mysql.GetConn().Debug()
	return nil
}

func (i *blogServiceImpl) Name() string {
	return blog.AppName
}
