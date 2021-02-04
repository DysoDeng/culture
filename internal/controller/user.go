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
	u := user.GetUserInfo(userId)
	if u.ID <= 0 {
		ctx.JSON(http.StatusOK, api.Fail(user.Error(), user.ErrorCode()))
		return
	}

	ctx.JSON(http.StatusOK, api.Success(u))
}
