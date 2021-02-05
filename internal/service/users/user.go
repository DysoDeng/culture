package users

import (
	"culture/internal/model"
	"culture/internal/service"
	"culture/internal/service/logs"
	"culture/internal/support/api"
	"culture/internal/support/db"
	"culture/internal/support/redis"
	"log"
)

type UserServiceInterface interface {
	service.Error
	GetUserInfo(userId int64) model.User
}

type UserService struct {
	service.BaseService
}

func NewUserService() *UserService {
	return &UserService{}
}

func (u *UserService) GetUserInfo(userId int64) model.User {
	if userId <= 0 {
		u.SetError("api.ErrorMissUid", api.CodeFail)
		return model.User{}
	}

	var user model.User
	db.DB().Debug().Where("id=?", userId).First(&user)
	if user.ID <= 0 {
		u.SetError("用户不存在", api.CodeFail)
		return model.User{}
	}

	res, _ := redis.Client().HGetAll("redis_key").Result()
	log.Println(res)

	var logService logs.LogServiceInterface
	_ = service.Provider(&logService)
	logService.Writer("aaaa")

	return user
}

type FinanceServiceInterface interface {
	GetUserFinance(userId int64)
}

type FinanceService struct{}

func NewFinanceService() *FinanceService {
	return &FinanceService{}
}

func (f FinanceService) GetUserFinance(userId int64) {
	log.Println(userId)
}
