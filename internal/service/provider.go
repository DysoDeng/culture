package service

import "github.com/goava/di"

// Container 服务容器
var Container *di.Container

// Provider 获取服务实例
func Provider(ptr di.Pointer, options ...di.ResolveOption) error {
	return Container.Resolve(ptr, options...)
}
