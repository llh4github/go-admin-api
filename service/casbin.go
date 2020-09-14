package service

import "github.com/llh4github/go-admin-api/vo"

// CasbinX Casbin服务层
type CasbinX struct {
}

// HasPermission 是否有权限
func HasPermission(rule vo.AuthRule) bool {
	// 放行不需要鉴权的请求
	isAnno, _ := enforcer.Enforce("anno", rule.URL, rule.Action)
	if isAnno {
		return true
	}
	for _, v := range rule.RoleNames {
		has, _ := enforcer.Enforce(v, rule.URL, rule.Action)
		if has {
			return true
		}
	}
	return false

}
