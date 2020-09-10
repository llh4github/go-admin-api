package vo

// JSONWrapper 统一响应结构体
type JSONWrapper struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// OkResponse 成功处理的响应
func OkResponse(data interface{}) JSONWrapper {
	return JSONWrapper{Code: 200, Msg: "ok", Data: data}
}

// ErrorResponse 异常响应
func ErrorResponse(code int, msg string) JSONWrapper {
	return JSONWrapper{Code: code, Msg: msg, Data: nil}
}

// ErrorResp 异常响应
func ErrorResp(info IRespErrorInfo) JSONWrapper {
	return JSONWrapper{Code: info.GetCode(), Msg: info.GetMsg(), Data: nil}
}

// ErrorRespWithData 异常响应（携带相关数据）
func ErrorRespWithData(info IRespErrorInfo, data interface{}) JSONWrapper {
	return JSONWrapper{Code: info.GetCode(), Msg: info.GetMsg(), Data: data}
}

// IRespErrorInfo 获取异常响应信息
type IRespErrorInfo interface {
	// GetCode 获取异常信息代码
	GetCode() int
	// GetMsg 获取异常信息文本消息
	GetMsg() string
}
