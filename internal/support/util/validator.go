package util

import "regexp"

// Validator 数据格式验证
type Validator struct{}

// IsMobile 判断是否为手机号
func (Validator) IsMobile(value string) bool {
	result, _ := regexp.MatchString(`^(1[0-9][0-9]\d{4,8})$`, value)
	return result
}

// IsPhone 判断是否为固定电话号码
func (Validator) IsPhone(value string) bool {
	result, _ := regexp.MatchString(`^(\d{4}-|\d{3}-)?(\d{8}|\d{7})$`, value)
	return result
}

// IsPhone400 判断是否为400电话
func (Validator) IsPhone400(value string) bool {
	result, _ := regexp.MatchString(`^400(-\d{3,4}){2}$`, value)
	return result
}

// IsEmail 判断是否为邮箱
func (Validator) IsEmail(value string) bool {
	result, _ := regexp.MatchString(`\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`, value)
	return result
}

// IsIdCard 是否身份证号
func (Validator) IsIdCard(value string) bool {
	result, _ := regexp.MatchString(`^\d{17}[0-9xX]$`, value)
	return result
}

// IsBankCard 是否银行卡号
func (Validator) IsBankCard(value string) bool {
	result, _ := regexp.MatchString(`^(\d{16}|\d{17}|\d{18}|\d{19})$`, value)
	return result
}
