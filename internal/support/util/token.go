package util

import (
	"culture/cloud/base/internal/config"
	"encoding/json"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

// token 数据结构
type Token struct {
	Token              json.Token `json:"token"`
	Expire             int64      `json:"expire"`
	RefreshToken       json.Token `json:"refresh_token"`
	RefreshTokenExpire int64      `json:"refresh_token_expire"`
}

// 生成用户Token
func GenerateToken(userType string, data map[string]interface{}) (Token, error) {

	currentTime := time.Now().Unix()
	var tokenMethod *jwt.Token
	var refreshTokenMethod *jwt.Token
	var expire int64
	var refreshTokenExpire int64

	conf := config.Config

	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()

	switch userType {
	case "user":
		expire = 24 * 3600
		refreshTokenExpire = 2 * 24 * 3600
		// Token
		tokenMethod = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"user_id":          data["user_id"],
			"user_type":        userType,
			"is_refresh_token": false,
			"iss":              conf.AppDomain + "/api/auth",
			"aud":              conf.AppDomain,
			"iat":              currentTime,
			"exp":              currentTime + int64(expire),
		})

		// RefreshToken
		refreshTokenMethod = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"user_id":          data["user_id"],
			"user_type":        userType,
			"is_refresh_token": true,
			"iss":              conf.AppDomain + "/api/auth",
			"aud":              conf.AppDomain,
			"iat":              currentTime,
			"exp":              currentTime + int64(expire),
		})
		break
	case "admin":
		expire = 4 * 3600
		refreshTokenExpire = 24 * 3600
		// Token
		tokenMethod = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"admin_id":         data["admin_id"],
			"user_type":        userType,
			"is_refresh_token": false,
			"iss":              conf.AppDomain + "/api/auth",
			"aud":              conf.AppDomain,
			"iat":              currentTime,
			"exp":              currentTime + int64(expire),
		})

		refreshTokenMethod = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"admin_id":         data["admin_id"],
			"user_type":        userType,
			"is_refresh_token": true,
			"iss":              conf.AppDomain + "/api/auth",
			"aud":              conf.AppDomain,
			"iat":              currentTime,
			"exp":              currentTime + int64(expire),
		})
	}

	if tokenMethod == nil {
		log.Println("tokenMethod nil")
		return Token{}, errors.New("token生成错误")
	}
	if refreshTokenMethod == nil {
		log.Println("refreshTokenMethod nil")
		return Token{}, errors.New("token生成错误")
	}

	// token
	var tokenSecret = []byte(config.Config.TokenKey)
	token, err := tokenMethod.SignedString(tokenSecret)
	if err != nil {
		return Token{}, errors.New("TOKEN生成错误")
	}

	// refreshToken
	refreshToken, err := refreshTokenMethod.SignedString(tokenSecret)
	if err != nil {
		return Token{}, errors.New("TOKEN生成错误")
	}

	return Token{
		Token:              token,
		Expire:             expire,
		RefreshToken:       refreshToken,
		RefreshTokenExpire: refreshTokenExpire,
	}, nil
}
