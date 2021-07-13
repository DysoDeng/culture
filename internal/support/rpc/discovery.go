package rpc

import (
	"context"
	"culture/cloud/base/internal/config"
	"fmt"
	"log"
	"sync"
	"time"

	rpcDiscovery "github.com/dysodeng/aux-rpc/discovery"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
)

var etcdV3Discovery *rpcDiscovery.EtcdV3Discovery

var services = make(map[string]*grpc.ClientConn)
var serviceLock sync.RWMutex

// Discovery 获取rpc服务连接
func Discovery(serviceName string, timeoutSecond int64) (context.Context, context.CancelFunc, *grpc.ClientConn) {

	// 连接超时
	if timeoutSecond <= 0 {
		timeoutSecond = 3
	}
	ctx, ctxCancel := context.WithDeadline(context.Background(), time.Now().Add(time.Duration(timeoutSecond)*time.Second))

	if c, ok := services[serviceName]; ok {
		return ctx, ctxCancel, c
	}

	// 连接rpc服务
	conn, err := grpc.Dial(
		fmt.Sprintf("%s:///%s", etcdV3Discovery.Scheme(), serviceName),
		grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, "round_robin")),
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)

	if err != nil {
		log.Fatalln(err)
	}

	serviceLock.Lock()
	services[serviceName] = conn
	serviceLock.Unlock()

	return ctx, ctxCancel, conn
}

func init() {
	etcdClient, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{config.Config.Etcd.Addr + ":" + config.Config.Etcd.Port},
		DialTimeout: 5 * time.Second,
	})

	if err != nil {
		log.Fatalln(err)
	}

	etcdV3Discovery = rpcDiscovery.NewEtcdV3Discovery(etcdClient, config.RpcPrefix)
}
