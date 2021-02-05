package middleware

import (
	"culture/internal/config"
	"culture/internal/support/api"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// Token验证
func TokenAuth(ctx *gin.Context) {
	tokenString := ctx.GetHeader("Authorization")

	conf := config.Config

	if tokenString == "" {
		ctx.Abort()
		ctx.JSON(http.StatusOK, api.Fail("miss token", api.CodeUnauthorized))
		return
	}

	token, err := jwt.Parse(tokenString, func(jwtToken *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", jwtToken.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(config.Config.TokenKey), nil
	})
	if err != nil {
		log.Println(err)
		ctx.Abort()
		ctx.JSON(http.StatusOK, api.Fail("token错误", api.CodeUnauthorized))
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if claims["aud"] != conf.AppDomain || claims["iss"] != conf.AppDomain+"/api/auth" {
			ctx.Abort()
			ctx.JSON(http.StatusOK, api.Fail("illegal token", api.CodeUnauthorized))
			return
		}
		if claims["is_refresh_token"] == true {
			ctx.Abort()
			ctx.JSON(http.StatusOK, api.Fail("refresh token不能用于业务请求", api.CodeUnauthorized))
			return
		}

		switch claims["user_type"].(string) {
		case "user":
			ctx.Set("user_type", "user")
			ctx.Set("user_id", int64(claims["user_id"].(float64)))
			break
		case "admin":
			ctx.Set("user_type", "admin")
			ctx.Set("admin_id", int64(claims["admin_id"].(float64)))
			break
		default:
			ctx.Abort()
			ctx.JSON(http.StatusOK, api.Fail("用户类型错误", api.CodeUnauthorized))
			return
		}

		ctx.Next()
	} else {
		ctx.Abort()
		ctx.JSON(http.StatusOK, api.Fail("illegal token", api.CodeUnauthorized))
		return
	}
}
