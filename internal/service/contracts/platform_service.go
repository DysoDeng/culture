package contracts

import "culture/cloud/base/internal/service"

// PlatformService 平台可定制服务接口
type PlatformService interface {
	// Name 获取服务名称
	Name() string
	// Ident 获取服务标识
	Ident() string
	// InitService 初始化服务
	InitService(cloudId uint64) service.Error
	// CheckService 检查服务是否可用
	CheckService(cloudId uint64) (bool, service.Error)
}
