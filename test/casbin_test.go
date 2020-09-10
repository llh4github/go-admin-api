package test

import (
	"testing"
)

// 测试数据（ACL模型）。按需执行
func TestInitData(t *testing.T) {
	e.AddPolicy("admin", "/api/admin/1", "post")
	e.AddPolicy("admin", "/api/admin/2", "get")
	e.AddPolicy("admin", "/api/admin/3", "put")
	e.AddPolicy("staff", "/api/staff/1", "get")
	e.AddPolicy("staff", "/api/staff/2", "post")
}

// 测试 权限验证

func TestValid(t *testing.T) {
	has1 := e.HasPermissionForUser("admin", "/api/admin/1", "post")
	has2 := e.HasPermissionForUser("admin", "/api/staff/1", "get")

	t.Logf("admin 有`%s`权限吗？ %v ", "/api/admin/1", has1)
	t.Logf("admin 有`%s`权限吗？ %v ", "/api/staff/1", has2)
	// ACL模型不能使用ROLE相关的API
	// 下面这句铁定报错
	// e.AddRoleForUser("tom", "admin")

}

/*
// 下面是RBAC模型的测试
// 测试了半天发现这个模型感觉并不好用还是ACL模型用着方便点
// 不同模型之间能使用的API是不同的，反正搞不好就报错

// 测试数据。按需执行
func TestInitData(t *testing.T) {

	e.AddPolicy("admin", "/api/admin/1", "post")
	e.AddPolicy("admin", "/api/admin/2", "get")
	e.AddPolicy("admin", "/api/admin/3", "put")
	e.AddPolicy("staff", "/api/staff/1", "get")
	e.AddPolicy("staff", "/api/staff/2", "post")

	e.AddRoleForUser("Tom", "admin")
	e.AddRoleForUser("Jerry", "staff")
}

// 测试 查询所有权限
func TestGetPolicy(t *testing.T) {
	p := e.GetPolicy()
	t.Log(p)
}

// 测试 角色添加与删除
func TestAddRole(t *testing.T) {

	t.Log("所有角色 ", e.GetAllRoles())
	add, err := e.AddRoleForUser("tom", "test")
	t.Log("给用户tom添加角色 ： ", add)
	del, err := e.DeleteRoleForUser("tom", "test")
	t.Log("删除用户tom的角色 ： ", del)
	t.Log(err)
}

// 测试 权限的添加与删除
func TestAddPermission(t *testing.T) {
	t.Log("所有权限 ", e.GetPolicy())
	add, err := e.AddPolicy("test", "/demo1", "get")
	// e.AddPolicy("test1", "/demo1", "get")
	t.Log("给 test 角色添加权限 ： ", add)
	// DeletePermission 删除所有指定的权限（因为一个权限可以分配给多个角色）
	// del, err := e.DeletePermission("/demo1", "get")
	// DeletePermissionForUser 删除指定角色下的权限。API名称设计得有点让人疑惑
	del, err := e.DeletePermissionForUser("test", "/demo1", "get")
	t.Log("删除 test 角色的一条权限 ", del)
	t.Log(err)
}

// 测试
func TestAuth(t *testing.T) {
	// e.AddPermissionForUser("Tom", "/api/admin/2", "get")
	p := e.GetPermissionsForUser("Tom")
	t.Log("Tom 所有权限", p)

	has := e.HasPermissionForUser("Tom", "/api/admin/2", "get")
	has2, _ := e.HasRoleForUser("Jerry", "admin")
	has3, _ := e.HasRoleForUser("Tom", "admin")
	has4 := e.HasPermissionForUser("Jerry", "/api/staff/2", "post")
	t.Log("Tom 有 '/api/staff/2' 权限吗？ ", has)
	t.Log("Jerry 是 admin 吗？", has2)
	t.Log("Tom 是 admin 吗？", has3)
	t.Log("Jerry 有 '/api/staff/1' 权限吗？", has4)
}
*/
