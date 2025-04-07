package common

import "time"

func NewMeta() *Meta {
	return &Meta{
		CreatedAt: time.Now().Unix(),
	}
}

type Meta struct {
	//用户id
	Id int `json:"id" gorm:"column:id"`
	//创建时间
	CreatedAt int64 `json:"created_at" gorm:"column:created_at"`
	//更新时间
	UpdatedAt int64 `json:"updated_at" gorm:"column:updated_at"`
}
