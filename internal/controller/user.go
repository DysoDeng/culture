package controller

import (
	userService "culture/internal/service/user"
	"culture/internal/util/api"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUser(ctx *gin.Context) {
	userId := ctx.MustGet("user_id").(int64)
	if userId <= 0 {
		ctx.JSON(http.StatusOK, api.Fail(api.ErrorMissUid, api.CodeFail))
		return
	}

	user := userService.NewUserService()
	u, err := user.GetUserInfo(userId)
	if err != nil {
		ctx.JSON(http.StatusOK, api.Fail(err.Error(), api.CodeFail))
		return
	}

	ctx.JSON(http.StatusOK, api.Success(u))
}
