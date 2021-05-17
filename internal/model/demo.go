package model

import (
	"culture/cloud/base/internal/config"
	"gorm.io/gorm"
)

// Demo 普通模型
type Demo struct {
	PrimaryKeyID        // 主键ID
	TestField    string `gorm:"type:varchar(150);not null;default:'';comment:测试字段" json:"test_field"`
	Time                // 添加时间，修改时间
}

func (Demo) TableName() string {
	return config.Config.DataBase.Prefix + "demo"
}

// DistributedDemo 分布式ID模型
type DistributedDemo struct {
	DistributedPrimaryKeyID
	OrderIndexKey
	TestField string `gorm:"type:varchar(150);not null;default:'';comment:测试字段" json:"test_field"`
	Time
}

func (DistributedDemo) TableName() string {
	return config.Config.DataBase.Prefix + "distributed_demo"
}

// BeforeCreate 创建分布式ID
func (distributedDemo *DistributedDemo) BeforeCreate(tx *gorm.DB) (err error) {
	distributedDemo.ID, distributedDemo.OrderIndex = CreateDistributedPrimaryKeyID()
	return
}
