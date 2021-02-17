package logs

import (
	"log"
)

// LogService 日志服务
type LogService struct{}

// NewLogService 初始化日志服务
func NewLogService() *LogService {
	return &LogService{}
}

// Writer 日志记录
func (l LogService) Writer(message string) {
	log.Println(message)
}
