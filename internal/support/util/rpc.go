package util

import (
	"context"
	"culture/cloud/base/internal/config"
	"log"
	"time"

	"github.com/dysodeng/drpc/discovery"
	"github.com/pkg/errors"
)

// RPCDiscovery 获取Rpc服务连接地址
func RPCDiscovery(timeoutSecond int64) (context.Context, context.CancelFunc, discovery.ServiceDiscovery, error) {
	d, err := discovery.NewEtcdV3Discovery([]string{config.Config.Etcd.Addr + ":" + config.Config.Etcd.Port}, config.RpcPrefix)
	if err != nil {
		log.Println(err)
		return nil, nil, nil, errors.Wrap(err, "rpc auth service error")
	}

	// 连接超时
	if timeoutSecond <= 0 {
		timeoutSecond = 3
	}
	ctx, ctxCancel := context.WithDeadline(context.Background(), time.Now().Add(time.Duration(timeoutSecond)*time.Second))

	return ctx, ctxCancel, d, nil
}
