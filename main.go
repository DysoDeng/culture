package main

import (
	"context"
	"culture/cloud/base/internal/config"
	"culture/cloud/base/internal/provider"
	"culture/cloud/base/internal/router"
	"culture/cloud/base/internal/support/util"
	"culture/cloud/base/rpc/proto"
	"culture/cloud/base/rpc/service"
	"github.com/dysodeng/drpc"
	"github.com/dysodeng/drpc/register"
	"github.com/gin-gonic/gin"
	"github.com/rcrowley/go-metrics"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {

	if config.Config.Env == config.Release {
		gin.SetMode(gin.ReleaseMode)
	}

	httpPort := "8080"
	rpcPort := "9000"

	// error logger
	logFilename := config.LogPath + "/gin.error.log"
	errLogFile, _ := os.OpenFile(logFilename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	gin.DefaultErrorWriter = io.MultiWriter(errLogFile, os.Stderr)

	// service container
	provider.ServiceProvider()

	ip := util.GetLocalIp()

	// grpc server
	rpcServer := drpc.NewServer(&register.EtcdV3Register{
		ServiceAddress: ip + ":" + rpcPort,
		EtcdServers:    []string{config.Config.Etcd.Addr + ":" + config.Config.Etcd.Port},
		BasePath:       config.RpcPrefix,
		Lease:          5,
		Metrics:        metrics.NewMeter(),
		ShowMetricsLog: true,
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

	// http server
	server := http.Server{
		Addr:    ":" + httpPort,
		Handler: router.Router(),
	}

	log.Printf("listening and serving HTTP on: :%s\n", httpPort)

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("http server error: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("shutdown http server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("http server shutdown error:", err)
	}
	log.Println("http server stop")
	log.Println("shutdown rpc server ...")
	_ = rpcServer.Stop()
}
