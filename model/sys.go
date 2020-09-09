package model

// ----- 系统基础模型 -------

// User 用户模型
type User struct {
	Base
	Username string `json:"username" gorm:"type:varchar(200)"`
	Password string `json:"password" gorm:"type:varchar(200)"`
}
