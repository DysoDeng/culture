package controller

import (
	"culture/cloud/base/internal/service"
	"culture/cloud/base/internal/service/contracts"
	"culture/cloud/base/internal/support/api"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Demo(ctx *gin.Context) {

	var demoService contracts.DemoServiceInterface
	_ = service.Resolve(&demoService)

	demo, err := demoService.Test("test")
	if err.Code != 200 {
		ctx.JSON(http.StatusOK, api.Fail(err.Error.Error(), err.Code))
		return
	}

	ctx.JSON(http.StatusOK, api.Success(demo.TestField))
}
