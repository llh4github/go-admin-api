package test

import (
	"testing"

	"github.com/llh4github/go-admin-api/utils"
)

// 测试加密和匹配
func TestHash(t *testing.T) {
	raw := "password"
	hashPwd, err := utils.HashPassword(raw)
	if err != nil {
		t.Log("密码加密失败！")
	}
	t.Log(hashPwd)
	match := utils.MatchPassword(raw, hashPwd)
	t.Logf("raw pwd :%s hashPwd: %s  密码匹配吗？ %v \n", raw, hashPwd, match)
	other := "other"
	match = utils.MatchPassword(other, hashPwd)
	t.Logf("raw pwd :%s hashPwd: %s  密码匹配吗？ %v \n", other, hashPwd, match)

}
