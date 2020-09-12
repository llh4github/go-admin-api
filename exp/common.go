// Package exp 异常信息包
package exp

// ExpCode 异常信息代码
type ExpCode int

const (
	// CommonError 系统出错
	CommonError ExpCode = 30000 + iota
	// BindJSONError 数据绑定错误
	BindJSONError
	// DataNotFound 数据不存在
	DataNotFound
)

// CommonExp 公共异常
type CommonExp struct {
	code int
	msg  string
}

// GetCode 获取异常信息代码
func (e CommonExp) GetCode() int {
	return e.code
}

// GetMsg 获取异常信息文本消息
func (e CommonExp) GetMsg() string {
	return e.msg
}

// GetCommonExp   通过异常代码获取公共异常的异常信息
func GetCommonExp(code ExpCode) CommonExp {
	_code := code
	_msg := ""
	switch code {
	case CommonError:
		_msg = "系统出错"
	case BindJSONError:
		_msg = "数据绑定错误"
	case DataNotFound:
		_msg = "数据不存在"
	default:
		_msg = "未知异常"

	}

	return CommonExp{code: int(_code), msg: _msg}
}
