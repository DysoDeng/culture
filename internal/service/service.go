package service

import "culture/internal/util/api"

// 服务错误接口
type Error interface {
	Error() string
	ErrorCode() api.Code
	SetError(error string, code api.Code)
}

// 基础服务
type BaseService struct {
	// 错误码
	code api.Code
	// 错误信息
	error string
}

// Error 获取错误信息
func (service *BaseService) Error() string {
	return service.error
}

// ErrorCode 获取错误码
func (service *BaseService) ErrorCode() api.Code {
	return service.code
}

// SetError 设置错误信息
// error 错误信息
// code 错误码
func (service *BaseService) SetError(error string, code api.Code) {
	if code <= 0 {
		code = api.CodeFail
	}
	service.error = error
	service.code = code
}