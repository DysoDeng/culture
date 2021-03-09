package router

import (
	"culture/cloud/base/internal/middleware"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {

	router := gin.Default()

	// 跨域处理
	router.Use(middleware.CrossDomain)

	api := router.Group("/v1/api")

	// 用户登录授权
	authRouter(api)

	// 鉴权路由
	authorization := api.Group("/")
	authorization.Use(middleware.TokenAuth)
	{
		userRouter(authorization)
	}

	return router
}
