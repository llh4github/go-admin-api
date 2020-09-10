package model

// ----- 系统基础模型 -------

// User 用户模型
type User struct {
	Base
	Username string `json:"username" gorm:"type:varchar(200)"`
	Password string `json:"password" gorm:"type:varchar(200)"`
}

// Role 角色
type Role struct {
	Base
	// RoleName 角色代号，全英文。用在系统内逻辑判断
	RoleName string `json:"role_name" gorm:"type:varchar(100)"`
	// DisplayName 显示名。给用户看的。
	DisplayName string `json:"display_name" gorm:"type:varchar(100)"`
	// Remark 备注
	Remark string `json:"remark" gorm:"type:varchar(200)"`
}
