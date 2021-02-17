package router

import (
	"culture/internal/controller"
	"github.com/gin-gonic/gin"
)

func authRouter(router *gin.RouterGroup) {
	apiAuth := router.Group("/auth")
	{
		apiAuth.POST("/login", controller.Login)
		apiAuth.POST("/refresh_token", controller.RefreshToken)
	}
}
