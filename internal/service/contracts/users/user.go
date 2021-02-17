package users

import (
	"culture/internal/model"
	"culture/internal/service"
)

// UserServiceInterface 用户服务接口
type UserServiceInterface interface {
	GetUserInfo(userId int64) (model.User, service.Error)
}

// FinanceServiceInterface 用户财务接口
type FinanceServiceInterface interface {
	GetUserFinance(userId int64) service.Error
}
