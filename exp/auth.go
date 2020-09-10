package exp

const (
	// AuthError 授权异常
	AuthError ExpCode = 40100 + iota
	// PwdError 用户或密码错误
	PwdError
)

// AuthExp 用户认证异常
type AuthExp struct {
	code int
	msg  string
}

// GetCode 获取异常信息代码
func (e AuthExp) GetCode() int {
	return e.code
}

// GetMsg 获取异常信息文本消息
func (e AuthExp) GetMsg() string {
	return e.msg
}

// GetAuthExp 通过异常代码获取用户认证相关的异常信息
func GetAuthExp(code ExpCode) AuthExp {
	_code := code
	_msg := ""
	switch code {
	case AuthError:
		_msg = "授权异常"
	case PwdError:
		_msg = "用户或密码错误"
	default:
		_msg = "未知异常"
	}
	return AuthExp{code: int(_code), msg: _msg}
}
