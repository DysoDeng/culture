package contracts

import (
	"culture/cloud/base/internal/model"
	"culture/cloud/base/internal/service"
)

// DemoServiceInterface 服务接口
type DemoServiceInterface interface {
	Test(params string) (model.Demo, service.Error)
}
