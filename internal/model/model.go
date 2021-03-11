package model

import "culture/cloud/base/internal/support/db"

// 基础模型
type ID struct {
	ID        uint64  		`gorm:"primary_key;autoIncrement" json:"id"`
}
type Time struct {
	CreatedAt db.JSONTime 	`gorm:"not null" json:"created_at"`
	UpdatedAt db.JSONTime	`gorm:"not null" json:"updated_at"`
}
