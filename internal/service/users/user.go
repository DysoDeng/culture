package users

import (
	"culture/internal/model"
	"culture/internal/service"
	"culture/internal/service/contracts/logs"
	"culture/internal/support/api"
	"culture/internal/support/db"
	"log"
)

// UserService 用户服务
type UserService struct{}

// NewUserService 初始化用户服务
func NewUserService() *UserService {
	return &UserService{}
}

// GetUserInfo 获取用户信息
func (u *UserService) GetUserInfo(userId int64) (model.User, service.Error) {
	if userId <= 0 {
		return model.User{}, service.Error{Code: api.CodeFail, Error: api.ErrorMissUid}
	}

	var user model.User
	db.DB().Debug().Where("id=?", userId).First(&user)
	if user.ID <= 0 {
		return model.User{}, service.Error{Code: api.CodeFail, Error: "用户不存在"}
	}

	//res, _ := redis.Client().HGetAll("redis_key").Result()
	//log.Println(res)

	var logService logs.LogServiceInterface
	_ = service.Resolve(&logService)
	logService.Writer("aaaa")

	return user, service.Error{Code: api.CodeOk}
}

// FinanceService 用户财务服务
type FinanceService struct{}

// NewFinanceService 初始化用户财务服务
func NewFinanceService() *FinanceService {
	return &FinanceService{}
}

// GetUserFinance 获取用户余额
func (f *FinanceService) GetUserFinance(userId int64) service.Error {
	log.Println(userId)
	return service.Error{Code: api.CodeOk}
}
