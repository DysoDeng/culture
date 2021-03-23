package model

import "culture/cloud/base/internal/support/db"

// 基础模型
// 所有数据模型都应该继承 PrimaryKeyID 与 Time 模型

// 主键ID
type PrimaryKeyID struct {
	ID uint64 `gorm:"primary_key;autoIncrement" json:"id"`
}

// 添加时间,修改时间
type Time struct {
	CreatedAt db.JSONTime `gorm:"type:datetime(0); not null" json:"created_at"`
	UpdatedAt db.JSONTime `gorm:"type:datetime(0); not null" json:"updated_at"`
}
