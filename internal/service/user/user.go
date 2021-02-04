package user

import (
	"culture/internal/model"
	"culture/internal/service"
	"culture/internal/support/db"
	"culture/internal/util/api"
)

type IUserService interface {
	service.Error
	GetUserInfo(userId int64) model.User
}

type User struct {
	service.BaseService
}

func NewUserService() IUserService {
	var userService IUserService = new(User)
	return userService
}

func (u *User) GetUserInfo(userId int64) model.User {
	if userId <= 0 {
		u.SetError("api.ErrorMissUid", api.CodeFail)
		return model.User{}
	}

	var user model.User
	db.DB.Debug().Where("id=?", userId).First(&user)
	if user.ID <= 0 {
		u.SetError("用户不存在", api.CodeFail)
		return model.User{}
	}

	return user
}
