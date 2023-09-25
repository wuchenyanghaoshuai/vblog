package common

import "time"


type Meta struct {

	Id int64 `json:"id"`
	//创建时间
	CreatedAt int64 `json:"created_at"`
	//更新时间
	UpdatedAt int64 `json:"updated_at"`

}

func NewMeta() *Meta{
	return &Meta{
		CreatedAt: time.Now().Unix(),
	}
}