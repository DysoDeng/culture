package middleware

import (
	"context"
	"culture/cloud/base/internal/config"
	"culture/cloud/base/internal/support/api"
	"culture/cloud/base/server/rpc/proto/auth"
	"github.com/dysodeng/drpc/discovery"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

// TokenAuth Token验证
func TokenAuth(ctx *gin.Context) {
	tokenString := ctx.GetHeader("Authorization")

	if tokenString == "" {
		ctx.Abort()
		ctx.JSON(http.StatusOK, api.Fail("miss token", api.CodeUnauthorized))
		return
	}

	d, err := discovery.NewEtcdV3Discovery([]string{config.Config.Etcd.Addr+":"+config.Config.Etcd.Port}, config.RpcPrefix)
	if err != nil {
		log.Println(err)
		ctx.Abort()
		ctx.JSON(http.StatusOK, api.Fail("rpc auth service error", api.CodeUnauthorized))
		return
	}
	defer d.Close()

	authCtx, authCancel := context.WithDeadline(context.Background(), time.Now().Add(3 * time.Second))
	defer authCancel()

	conn := d.Conn("AuthService")
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
		d, err := discovery.NewEtcdV3Discovery([]string{config.Config.Etcd.Addr+":"+config.Config.Etcd.Port}, config.RpcPrefix)
		if err != nil {
			log.Println(err)
			ctx.Abort()
			ctx.JSON(http.StatusOK, api.Fail("rpc auth service error", api.CodeUnauthorized))
			return
		}
		defer d.Close()

		authCtx, authCancel := context.WithDeadline(context.Background(), time.Now().Add(3 * time.Second))
		defer authCancel()

		conn := d.Conn("AuthService")
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
