package router

import (
	"culture/cloud/base/internal/controller"
	"github.com/gin-gonic/gin"
)

func userRouter(router *gin.RouterGroup) {
	userRoute := router.Group("/user")
	{
		userRoute.POST("/info", controller.GetUser)
		userRoute.POST("/finance", controller.GetFinance)
	}
}
