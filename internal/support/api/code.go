package api

// 响应码
type Code int

// 接口响应码
const (
	CodeOk               Code = 200 // 成功
	CodeFail             Code = 0   // 业务错误
	CodeBadRequest       Code = 400 // 请求出错
	CodeUnauthorized     Code = 401 // 未授权
	CodeForbidden        Code = 403 // 无权限
	CodeNotFound         Code = 404 // 未找到
	CodeMethodNotAllowed Code = 405 // 请求方法不允许

	// 错误消息内容
	ErrorBusy    string = "系统繁忙，请稍后再试"
	ErrorMissUid string = "缺少用户ID"
)
