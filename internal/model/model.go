package model

import (
	"culture/cloud/base/internal/support/db"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"strings"
	"time"
)

// 基础模型
// 所有数据模型都应该继承 PrimaryKeyID 或 DistributedPrimaryKeyID 与 Time 模型

// PrimaryKeyID 自增主键ID
type PrimaryKeyID struct {
	ID uint64 `gorm:"primary_key;autoIncrement" json:"id"`
}

// DistributedPrimaryKeyID 分布式主键ID
type DistributedPrimaryKeyID struct {
	ID string `gorm:"type:varchar(64);primary_key" json:"id"`
}

// OrderIndexKey 排序索引
type OrderIndexKey struct {
	OrderIndex uint64 `gorm:"index;not null;default 0;comment:排序索引" json:"order_index"`
}

// Time 添加时间,修改时间
type Time struct {
	CreatedAt db.JSONTime `gorm:"type:datetime(0);index;not null" json:"created_at,omitempty"`
	UpdatedAt db.JSONTime `gorm:"type:datetime(0);not null" json:"updated_at,omitempty"`
}

// CreateDistributedPrimaryKeyID 创建分布式ID
func CreateDistributedPrimaryKeyID() (id string, orderIndex uint64) {
	return strings.Replace(uuid.New().String(), "-", "", -1), uint64(time.Now().UnixNano())
}

// IncrementInt64 整数自增
func IncrementInt64(column string, amount int64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		db.UpdateColumn(column, gorm.Expr(column+" + ?", amount))
		return db
	}
}

// IncrementFloat64 浮点数自增
func IncrementFloat64(column string, amount float64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		db.UpdateColumn(column, gorm.Expr(column+" + ?", amount))
		return db
	}
}

// DecrementInt64 整数自减
func DecrementInt64(column string, amount int64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		db.UpdateColumn(column, gorm.Expr(column+" - ?", amount))
		return db
	}
}

// DecrementFloat64 浮点数自减
func DecrementFloat64(column string, amount float64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		db.UpdateColumn(column, gorm.Expr(column+" - ?", amount))
		return db
	}
}
