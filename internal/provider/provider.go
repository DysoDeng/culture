package provider

import (
	"culture/internal/service"
	"culture/internal/service/logs"
	"culture/internal/service/users"
	"github.com/goava/di"
)

// 服务容器初始化
func ServiceProvider() {
	di.SetTracer(&di.StdTracer{})
	var err error
	service.Container, err = di.New(
		di.Provide(users.NewUserService, di.As(new(users.UserServiceInterface))),
		di.Provide(users.NewFinanceService, di.As(new(users.FinanceServiceInterface))),
		di.Provide(logs.NewLogService, di.As(new(logs.LogServiceInterface))),
	)
	if err != nil {
		panic(err)
	}
}
