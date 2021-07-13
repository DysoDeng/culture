package rpc

import (
	"culture/cloud/base/internal/config"
	"culture/cloud/base/server/rpc/proto"
	"culture/cloud/base/server/rpc/service"

	auxrpc "github.com/dysodeng/aux-rpc"
	"github.com/dysodeng/aux-rpc/registry"
	"github.com/rcrowley/go-metrics"
)

// Server grpc server
func Server(ip string, rpcPort string) *auxrpc.Server {
	// grpc server
	rpcServer := auxrpc.NewServer(&registry.EtcdV3Registry{
		ServiceAddress: ip + ":" + rpcPort,
		EtcdServers:    []string{config.Config.Etcd.Addr + ":" + config.Config.Etcd.Port},
		BasePath:       config.RpcPrefix,
		Lease:          5,
		Metrics:        metrics.NewMeter(),
		ShowMetricsLog: false,
	})
	defer func() {
		if err := recover(); err != nil {
			_ = rpcServer.Stop()
		}
	}()

	_ = rpcServer.Register("DemoService", &service.DemoService{}, proto.RegisterDemoServer, "")

	go func() {
		rpcServer.Serve(":" + rpcPort)
	}()

	return rpcServer
}
