package service

import (
	"github.com/llh4github/go-admin-api/vo"
)

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

// Add a rule
func (c CasbinX) Add(info vo.PermInfo) bool {
	add, _ := enforcer.AddPermissionForUser(info.RoleName, info.URL, info.Action)
	return add

}

// All 查询所有
func (c CasbinX) All() []vo.PermInfo {

	var list []vo.PermInfo
	db.Raw(`
	SELECT
	casbin_rule.v0 AS role_name,
	casbin_rule.v1 AS url,
	casbin_rule.v2 AS action
	FROM
	casbin_rule
	`).Scan(&list)
	return list
}

// Delete 删除信息
func (c CasbinX) Delete(info vo.PermInfo) bool {
	deleted, _ := enforcer.DeletePermissionForUser(info.RoleName, info.URL, info.Action)
	return deleted
}
