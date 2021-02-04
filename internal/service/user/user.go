package user

import (
	"culture/internal/model"
	"culture/internal/support/db"
	"culture/internal/support/redis"
	"culture/internal/util/api"
	"errors"
	"log"
)

type IUserService interface {
	GetUserInfo(userId int64) (model.User, error)
}

type User struct {
}

func NewUserService() IUserService {
	var userService IUserService = new(User)
	return userService
}

func (u User) GetUserInfo(userId int64) (model.User, error) {
	if userId <= 0 {
		return model.User{}, errors.New(api.ErrorMissUid)
	}

	var user model.User
	db.DB.Debug().Where("id=?", userId).First(&user)
	if user.ID <= 0 {
		return model.User{}, errors.New("用户不存在")
	}

	cache := redis.Client()
	b, err := cache.HMSet("redis_key", map[string]interface{}{"key": "value"}).Result()
	log.Println(b, err)
	d, _ := cache.HGetAll("redis_key").Result()
	log.Println(d)

	return user, nil
}
