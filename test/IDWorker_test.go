package test

import (
	"testing"

	"github.com/llh4github/go-admin-api/utils"
)

// 测试 id生成方法
func TestID(t *testing.T) {

	t.Log(utils.NextIDDefalut())
	t.Log(utils.NextIDDefalut())
	t.Log(utils.NextIDDefalut())
	t.Log(utils.NextIDDefalut())
}
