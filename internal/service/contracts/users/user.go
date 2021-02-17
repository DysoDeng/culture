package users

import (
	"culture/internal/model"
	"culture/internal/service"
)

type UserServiceInterface interface {
	service.Error
	GetUserInfo(userId int64) model.User
}

type FinanceServiceInterface interface {
	GetUserFinance(userId int64)
}
