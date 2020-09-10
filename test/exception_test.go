package test

import (
	"testing"

	"github.com/llh4github/go-admin-api/exp"
	"github.com/llh4github/go-admin-api/vo"
)

// 测试 通过异常代码获取异常信息
func TestExpInfo(t *testing.T) {
	e := exp.GetAuthExp(exp.PwdError)
	j := vo.ErrorResp(e)
	t.Log(j)
}
