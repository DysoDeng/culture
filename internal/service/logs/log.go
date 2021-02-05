package logs

import (
	"culture/internal/service"
	"log"
)

type LogServiceInterface interface {
	service.Error
	Writer(message string)
}

type LogService struct {
	service.BaseService
}

func NewLogService() *LogService {
	return &LogService{}
}

func (l *LogService) Writer(message string) {
	log.Println(message)
}