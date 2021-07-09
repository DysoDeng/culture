package message

import (
	"strings"

	"github.com/pkg/errors"
)

// ErrorMessage 错误信息类型
type ErrorMessage string

// String
func (em ErrorMessage) String() string {
	return string(em)
}

// Error
func (em ErrorMessage) Error() error {
	return errors.New(em.String())
}

// Param 参数设置
func (em ErrorMessage) Param(param map[string]string) ErrorMessage {
	str := em.String()
	for k, v := range param {
		str = strings.Replace(str, "{"+k+"}", v, 1)
	}
	return ErrorMessage(str)
}

// 错误信息
const (
	EMMissUsername        ErrorMessage = "登录账号为空，请输入邮箱或手机号码。"
	EMMissPassword        ErrorMessage = "登录密码为空，请输入至少6位字符(包含字母加数字或符号的组合)。"
	EMUsernameFormatError ErrorMessage = "你输入的账号格式不正确，请输入邮箱或手机号码。"
	EMPasswordFormatError ErrorMessage = "你输入的密码格式不正确，请输入至少6位字符(包含字母加数字或符号的组合)。"
	EMNonAgreementClause  ErrorMessage = "你需要同意《相关协议和条款》方能完成账号的注册和平台的使用。"
)
