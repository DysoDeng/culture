package http

import (
	"context"
	"culture/cloud/base/server/http/router"
	"log"
	"net/http"
	"time"
)

// Server http server
func Server(httpPort string) *http.Server {
	// http server
	server := http.Server{
		Addr:    ":" + httpPort,
		Handler: router.Router(),
	}

	defer func() {
		if err := recover(); err != nil {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			if err := server.Shutdown(ctx); err != nil {
				log.Fatal("http server shutdown error:", err)
			}
		}
	}()

	log.Printf("listening and serving HTTP on: :%s\n", httpPort)

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("http server error: %s\n", err)
		}
	}()

	return &server
}
