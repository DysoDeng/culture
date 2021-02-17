package provider

import (
	"culture/internal/service"
	"culture/internal/service/contracts/logs"
	"culture/internal/service/contracts/users"
	logsService "culture/internal/service/logs"
	usersService "culture/internal/service/users"
	"github.com/goava/di"
)

// 服务容器初始化
func ServiceProvider() {
	di.SetTracer(&di.StdTracer{})
	var err error
	service.Container, err = di.New(
		di.Provide(usersService.NewUserService, di.As(new(users.UserServiceInterface))),
		di.Provide(usersService.NewFinanceService, di.As(new(users.FinanceServiceInterface))),
		di.Provide(logsService.NewLogService, di.As(new(logs.LogServiceInterface))),
	)
	if err != nil {
		panic(err)
	}
}
