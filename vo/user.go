package vo

// Account 帐户
// 	用于登录、注册
type Account struct {
	Username, Password string
}

// UserRoles 用户-角色关系
type UserRoles struct {
	UserID  string   `json:"user_id"`
	RoleIDs []string `json:"role_ids"`
}

// AuthRule 认证规则
type AuthRule struct {
	RoleNames []string
	URL       string
	Action    string
}
