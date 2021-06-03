package contracts

import "culture/cloud/base/internal/service"

// App 平台应用接口
type App interface {
	// Name 获取应用名称
	Name() string
	// Ident 获取应用标识
	Ident() string
	// InitApp 初始化应用
	InitApp(cloudID uint64) service.Error
	// CheckApp 检查应用是否可用
	CheckApp(cloudID uint64) (bool, service.Error)
}
