package main

import (
	"context"
	"culture/cloud/base/internal/config"
	"culture/cloud/base/internal/provider"
	"culture/cloud/base/internal/support/util"
	"culture/cloud/base/server/http"
	"culture/cloud/base/server/rpc"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
	"os/signal"
	"time"
)

func main() {

	if config.Config.Env == config.Release {
		gin.SetMode(gin.ReleaseMode)
	}

	httpPort := config.Config.HttpPort
	rpcPort := config.Config.RpcPort
	if httpPort == "" {
		httpPort = "8080"
	}
	if rpcPort == "" {
		rpcPort = "9000"
	}

	// error logger
	logFilename := config.LogPath + "/gin.error.log"
	errLogFile, _ := os.OpenFile(logFilename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	gin.DefaultErrorWriter = io.MultiWriter(errLogFile, os.Stderr)

	// service container
	provider.ServiceProvider()

	ip := util.GetLocalIp()

	// grpc server
	rpcServer := rpc.Server(ip, rpcPort)

	// http server
	httpServer := http.Server(httpPort)

	// 等待中断信号以优雅地关闭服务器
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("shutdown http server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := httpServer.Shutdown(ctx); err != nil {
		log.Fatal("http server shutdown error:", err)
	}
	log.Println("http server stop")
	log.Println("shutdown rpc server ...")
	_ = rpcServer.Stop()
}
