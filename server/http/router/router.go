package router

import (
	"culture/cloud/base/server/http/controller"
	"culture/cloud/base/server/http/middleware"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {

	router := gin.Default()

	// 跨域处理
	router.Use(middleware.CrossDomain)

	api := router.Group("/api")

	api.POST("/demo", controller.Demo)

	return router
}
