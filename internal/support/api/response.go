package api

// Response api 响应数据结构
type Response struct {
	// Code 错误码
	Code Code `json:"code"`
	// Data data payload
	Data interface{} `json:"data,omitempty"`
	// Error 错误信息
	Error string `json:"error"`
}

// Success 正确响应
func Success(result interface{}) Response {
	return Response{CodeOk, result, "success"}
}

// Fail 失败响应
func Fail(error string, code Code) Response {
	return Response{code, nil, error}
}
