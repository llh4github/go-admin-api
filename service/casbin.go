package service

import "github.com/llh4github/go-admin-api/vo"

// CasbinX Casbin服务层
type CasbinX struct {
}

// HasPermission 是否有权限
func HasPermission(rule vo.AuthRule) bool {
	for _, v := range rule.RoleNames {
		has := enforcer.HasPermissionForUser(v, rule.URL, rule.Action)
		if has {
			return true
		}
	}
	return false

}
