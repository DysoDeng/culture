package model

import (
	"culture/cloud/base/internal/config"
	"culture/cloud/base/internal/support/db"
)

// 用户
type User struct {
	ID        int64       `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Telephone string      `gorm:"type:varchar(150);unique_index" json:"telephone"`
	RealName  string      `gorm:"default null" json:"real_name"`
	CreatedAt db.JSONTime `gorm:"not null" json:"created_at"`
}

func (User) TableName() string {
	return config.Config.DataBase.Prefix + "users"
}
