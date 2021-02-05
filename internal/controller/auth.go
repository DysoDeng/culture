package controller

import (
	"culture/internal/config"
	"culture/internal/model"
	"culture/internal/support/db"
	"culture/internal/util"
	"culture/internal/util/api"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type LoginAuth struct {
	UserType string `form:"user_type" json:"user_type" binding:"required"`
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type RegisterData struct {
	UserType        string `form:"user_type" json:"user_type" binding:"required"`
	Username        string `form:"username" json:"username" binding:"required"`
	Password        string `form:"password" json:"password" binding:"required"`
	ConfirmPassword string `form:"confirm_password" json:"confirm_password" binding:"required"`
}

// 用户登录
func Login(ctx *gin.Context) {

	var auth LoginAuth

	if ctx.ShouldBind(&auth) != nil {
		if auth.UserType == "" {
			ctx.JSON(http.StatusOK, api.Fail("用户类型错误", api.CodeFail))
			return
		}
		if auth.Username == "" {
			ctx.JSON(http.StatusOK, api.Fail("用户名为空", api.CodeFail))
			return
		}
		if auth.Password == "" {
			ctx.JSON(http.StatusOK, api.Fail("密码为空", api.CodeFail))
			return
		}
	}

	data := make(map[string]interface{})

	switch auth.UserType {
	case "user":
		var user model.User
		db.DB().Debug().Where("telephone=?", auth.Username).First(&user)
		if user.ID <= 0 {
			ctx.JSON(http.StatusOK, api.Fail("用户名错误", api.CodeFail))
			return
		}

		data["user_id"] = user.ID

		break
	default:
		ctx.JSON(http.StatusOK, api.Fail("用户类型错误", api.CodeFail))
		return
	}

	token, err := util.GenerateToken(auth.UserType, data)
	if err != nil {
		ctx.JSON(http.StatusOK, api.Fail(err.Error(), api.CodeFail))
		return
	}

	ctx.JSON(http.StatusOK, api.Success(token))
}

// Token刷新
func RefreshToken(ctx *gin.Context) {

	conf := config.Config

	refreshToken := ctx.PostForm("refresh_token")
	if refreshToken == "" {
		ctx.JSON(http.StatusOK, api.Fail("TOKEN刷新令牌未指定", api.CodeFail))
		return
	}

	token, err := jwt.Parse(refreshToken, func(jwtToken *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", jwtToken.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(config.Config.TokenKey), nil
	})
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusOK, api.Fail("token错误", api.CodeUnauthorized))
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if claims["aud"] != conf.AppDomain || claims["iss"] != conf.AppDomain+"/api/auth" {
			ctx.JSON(http.StatusOK, api.Fail("illegal token", api.CodeUnauthorized))
			return
		}
		if claims["is_refresh_token"] == false {
			ctx.JSON(http.StatusOK, api.Fail("业务token不能用于token刷新", api.CodeUnauthorized))
			return
		}

		data := make(map[string]interface{})

		userType := claims["user_type"].(string)

		switch userType {
		case "user":
			userId := int64(claims["user_id"].(float64))
			var user model.User
			db.DB().Debug().Table(db.FullTableName("users")).
				Where("id=?", userId).First(&user)
			if user.ID <= 0 {
				ctx.JSON(http.StatusOK, api.Fail("用户不存在", api.CodeFail))
				return
			}

			data["user_id"] = userId
			break
		default:
			ctx.JSON(http.StatusOK, api.Fail("用户类型错误", api.CodeUnauthorized))
			return
		}

		token, err := util.GenerateToken(userType, data)
		if err != nil {
			ctx.JSON(http.StatusOK, api.Fail(err.Error(), api.CodeFail))
			return
		}

		ctx.JSON(http.StatusOK, api.Success(token))
	} else {
		ctx.JSON(http.StatusOK, api.Fail("illegal token", api.CodeUnauthorized))
		return
	}
}
