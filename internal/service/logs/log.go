package logs

import (
	"culture/internal/service"
	"culture/internal/service/contracts/users"
	"log"
)

type LogService struct {
	service.BaseService
}

func NewLogService() *LogService {
	return &LogService{}
}

func (l *LogService) Writer(message string) {
	var userService users.UserServiceInterface
	_ = service.Provider(&userService)
	log.Println(userService.Error())
	log.Println(message)
}