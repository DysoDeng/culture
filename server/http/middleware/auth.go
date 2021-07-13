package middleware

import (
	"culture/cloud/base/internal/support/api"
	"culture/cloud/base/internal/support/rpc"
	"culture/cloud/base/server/rpc/proto/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

// TokenAuth Token验证
func TokenAuth(ctx *gin.Context) {
	tokenString := ctx.GetHeader("Authorization")

	if tokenString == "" {
		ctx.Abort()
		ctx.JSON(http.StatusOK, api.Fail("miss token", api.CodeUnauthorized))
		return
	}

	rpcCtx, rpcCancel, conn := rpc.Discovery("Passport/AuthService", 3)
	defer rpcCancel()

	authService := auth.NewAuthClient(conn)
	res, err := authService.ValidToken(rpcCtx, &auth.TokenRequest{Token: tokenString})
	if err != nil {
		ctx.Abort()
		ctx.JSON(http.StatusOK, api.Fail(err.Error(), api.CodeUnauthorized))
		return
	}
	if res.Code != api.CodeOk.ToInt64() {
		if res.Code == api.CodeNotCreate.ToInt64() {
			if res.UserType == "culture" {
				ctx.Abort()
				ctx.JSON(http.StatusOK, api.Fail("未创建文化云", api.CodeNotCreate))
				return
			}
		}
	} else {
		switch res.UserType {
		case "culture":
			ctx.Set("user_type", "culture")
			ctx.Set("master_id", res.UserId)
			ctx.Set("cloud_id", res.CloudId)
			break
		case "user":
			ctx.Set("user_type", "user")
			ctx.Set("user_id", res.UserId)
			break
		case "admin":
			ctx.Set("user_type", "admin")
			ctx.Set("admin_id", res.UserId)
			break
		default:
			ctx.Abort()
			ctx.JSON(http.StatusOK, api.Fail("用户类型错误", api.CodeUnauthorized))
			return
		}
	}

	ctx.Next()
}

// NotTokenAuth 通用Token验证
func NotTokenAuth(ctx *gin.Context) {
	tokenString := ctx.GetHeader("Authorization")

	if tokenString == "" {
		ctx.Next()
	} else {
		rpcCtx, rpcCancel, conn := rpc.Discovery("Passport/AuthService", 3)
		defer rpcCancel()

		authService := auth.NewAuthClient(conn)
		res, err := authService.ValidNotToken(rpcCtx, &auth.TokenRequest{Token: tokenString})
		if err != nil {
			ctx.Next()
		} else {
			if res.Code != api.CodeOk.ToInt64() {
				ctx.Next()
			} else {
				switch res.UserType {
				case "culture":
					ctx.Set("user_type", "culture")
					ctx.Set("master_id", res.UserId)
					ctx.Set("cloud_id", res.CloudId)
					break
				case "user":
					ctx.Set("user_type", "user")
					ctx.Set("user_id", res.UserId)
					break
				case "admin":
					ctx.Set("user_type", "admin")
					ctx.Set("admin_id", res.UserId)
					break
				default:
					ctx.Next()
				}
			}

			ctx.Next()
		}
	}
}
