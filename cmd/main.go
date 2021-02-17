package main

import (
	"context"
	"culture/internal/config"
	"culture/internal/provider"
	"culture/internal/router"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {

	if config.Config.Env == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	port := "8080"

	// error logger
	errLogFile, _ := os.Create("storage/logs/gin.error.log")
	gin.DefaultErrorWriter = io.MultiWriter(errLogFile, os.Stderr)

	// service container
	provider.ServiceProvider()

	server := http.Server{
		Addr: ":"+port,
		Handler: router.Router(),
	}

	log.Printf("listening and serving HTTP on: %s\n", port)

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

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("http server shutdown error:", err)
	}
	log.Println("http server stop")
}
