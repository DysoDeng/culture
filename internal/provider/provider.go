package provider

import (
	"culture/cloud/base/internal/config"
	"culture/cloud/base/internal/service"
	"culture/cloud/base/internal/service/contracts"
	"fmt"
	"github.com/goava/di"
	"github.com/pkg/errors"
)

// 服务容器初始化
func ServiceProvider() {
	if config.Config.Env != config.Release {
		di.SetTracer(&di.StdTracer{})
	}
	var err error
	service.Container, err = di.New(
		di.Provide(service.NewDemoService, di.As(new(contracts.DemoServiceInterface))),
	)
	if err != nil {
		err = errors.Wrap(err, "service provider error.")
		fmt.Printf("%+v\n", err)
		panic(err)
	}
}
