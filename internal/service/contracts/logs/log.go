package logs

// LogServiceInterface 日志服务接口
type LogServiceInterface interface {
	// Writer 日志记录
	Writer(message string)
}
