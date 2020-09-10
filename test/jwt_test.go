package test

import (
	"testing"

	"github.com/llh4github/go-admin-api/utils"
)

// 测试 token生成方法
func TestGenToken(t *testing.T) {
	token, err := utils.CreateToken("aaa")
	if err != nil {
		t.Log(err)
	}
	t.Log(token)
}

// 测试 token解析方法
func TestParse(t *testing.T) {
	token, err := utils.CreateToken("aaa")
	if err != nil {
		t.Log(err)
	}
	m, rs := utils.ParseToken(token)
	if !rs {
		t.Log("token 过期或不合法")
	}
	t.Log(m)

	sub, has := utils.GetSub(token)
	if !has {
		t.Log("token中不包含 sub 字段信息")
	}
	t.Log(sub)
	exp := utils.GetExp(token)
	t.Logf("token 将于 %d 过期", exp)
}
