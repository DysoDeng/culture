package contracts

import "culture/cloud/base/internal/service"

// Module 模块接口
type Module interface {
	// Name 获取模块名称
	Name() string
	// Ident 获取模块标识
	Ident() string
	// InitModule 初始化模块
	InitModule(cloudID uint64) service.Error
	// CheckModule 检查模块是否可用
	CheckModule(cloudID uint64) (bool, service.Error)
}
