package controller

import (
	"wuchenyanghaoshuai/vblog/skills/mvc/dao"
	"wuchenyanghaoshuai/vblog/skills/mvc/model"
)

func CreateUser(ins *model.User) error {
	dao.SaveUser(ins)
	return nil
}
