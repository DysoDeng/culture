package rpc

import (
	"culture/cloud/base/internal/config"
	"culture/cloud/base/server/rpc/proto"
	"culture/cloud/base/server/rpc/service"
	"github.com/dysodeng/drpc"
	"github.com/dysodeng/drpc/register"
	"github.com/rcrowley/go-metrics"
)

// Server grpc server
func Server(ip string, rpcPort string) *drpc.Server {
	// grpc server
	rpcServer := drpc.NewServer(&register.EtcdV3Register{
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

	_ = rpcServer.Register(&service.DemoService{}, proto.RegisterDemoServer, "")

	go func() {
		rpcServer.Serve(":" + rpcPort)
	}()

	return rpcServer
}
