package middleware

import (
	"culture/cloud/base/internal/support/api"
	"culture/cloud/base/internal/support/util"
	"culture/cloud/base/server/rpc/proto/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

// TokenAuth Token验证
func TokenAuth(ctx *gin.Context) {
	tokenString := ctx.GetHeader("Authorization")

	if tokenString == "" {
		ctx.Abort()
		ctx.JSON(http.StatusOK, api.Fail("miss token", api.CodeUnauthorized))
		return
	}

	authCtx, authCancel, rpcDiscovery, err := util.RpcDiscovery(3)
	if err != nil {
		ctx.JSON(http.StatusOK, api.Fail(err.Error(), api.CodeFail))
		return
	}
	defer rpcDiscovery.Close()
	defer authCancel()

	conn := rpcDiscovery.Conn("AuthService")
	authService := auth.NewAuthClient(conn)
	res, err := authService.ValidToken(authCtx, &auth.TokenRequest{Token: tokenString})
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
		authCtx, authCancel, rpcDiscovery, err := util.RpcDiscovery(3)
		if err != nil {
			ctx.JSON(http.StatusOK, api.Fail(err.Error(), api.CodeFail))
			return
		}
		defer rpcDiscovery.Close()
		defer authCancel()

		conn := rpcDiscovery.Conn("AuthService")
		authService := auth.NewAuthClient(conn)
		res, err := authService.ValidNotToken(authCtx, &auth.TokenRequest{Token: tokenString})
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
