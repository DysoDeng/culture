package controller

import (
	"culture/internal/service"
	"culture/internal/service/users"
	"culture/internal/support/api"
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
	_ = service.Provider(&userService)
	u := userService.GetUserInfo(userId)
	if u.ID <= 0 {
		ctx.JSON(http.StatusOK, api.Fail(userService.Error(), userService.ErrorCode()))
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
	_ = service.Provider(&financeService)
	financeService.GetUserFinance(userId)

	ctx.JSON(http.StatusOK, api.Success("1"))
}
