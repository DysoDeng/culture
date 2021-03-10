package model

import (
	"culture/cloud/base/internal/config"
	"culture/cloud/base/internal/support/db"
)

// 用户
type User struct {
	ID        uint64    `gorm:"primary_key;autoIncrement" json:"id"`
	Telephone string    `gorm:"type:varchar(150);uniqueIndex;not null;default:'';comment:手机号" json:"telephone"`
	RealName  string    `gorm:"type:varchar(150);not null;default:'';comment:真实姓名" json:"real_name"`
	Nickname  string    `gorm:"type:varchar(150);not null;default:'';comment:昵称" json:"nickname"`
	CreatedAt db.JSONTime 		`gorm:"not null" json:"created_at"`
}

func (User) TableName() string {
	return config.Config.DataBase.Prefix + "users"
}
