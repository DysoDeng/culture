package controller

import (
	"culture/cloud/base/internal/service"
	"culture/cloud/base/internal/service/contracts/users"
	"culture/cloud/base/internal/support/api"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUser(ctx *gin.Context) {
	userId := ctx.MustGet("user_id").(int64)
	if userId <= 0 {
		ctx.JSON(http.StatusOK, api.Fail(api.ErrorMissUid, api.CodeFail))
		return
	}

	var userService users.UserServiceInterface
	_ = service.Resolve(&userService)

	u, err := userService.GetUserInfo(userId)
	if err.Code != api.CodeOk {
		ctx.JSON(http.StatusOK, api.Fail(err.Error, err.Code))
		return
	}

	ctx.JSON(http.StatusOK, api.Success(u))
}

func GetFinance(ctx *gin.Context) {
	userId := ctx.MustGet("user_id").(int64)
	if userId <= 0 {
		ctx.JSON(http.StatusOK, api.Fail(api.ErrorMissUid, api.CodeFail))
		return
	}

	var financeService users.FinanceServiceInterface
	_ = service.Resolve(&financeService)
	financeService.GetUserFinance(userId)

	ctx.JSON(http.StatusOK, api.Success("1"))
}
