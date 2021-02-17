package api

// api 响应数据结构
type Response struct {
	Code  Code        `json:"code"`
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error"`
}

// 正确响应
func Success(result interface{}) Response {
	return Response{CodeOk, result, "success"}
}

// 失败响应
func Fail(error string, code Code) Response {
	return Response{code, nil, error}
}
