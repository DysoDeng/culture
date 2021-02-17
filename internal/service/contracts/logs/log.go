package logs

import "culture/internal/service"

type LogServiceInterface interface {
	service.Error
	Writer(message string)
}
