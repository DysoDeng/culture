package model

import (
	"culture/cloud/base/internal/config"
)

// Demo 模型
type Demo struct {
	ID // 主键ID
	TestField string    `gorm:"type:varchar(150);not null;default:'';comment:测试字段" json:"test_field"`
	Time // 添加时间，修改时间
}
func (Demo) TableName() string {
	return config.Config.DataBase.Prefix + "demo"
}
