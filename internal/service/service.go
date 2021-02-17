package service

import (
	"culture/internal/support/api"
	"github.com/goava/di"
)

// Error 服务错误码
type Error struct {
	Code  api.Code
	Error string
}

// Container 服务容器
var Container *di.Container

// Resolve 获取服务实例
func Resolve(ptr di.Pointer, options ...di.ResolveOption) error {
	return Container.Resolve(ptr, options...)
}
