package service

import "github.com/llh4github/go-admin-api/model"

// User 用户信息服务层
type User struct {
}

// UniqueUsername 检查用户名是否已存在
// 	true 输入用户名是唯一的、可用
func UniqueUsername(username string) bool {
	var user model.User
	result := db.Where("username = ? and remove_flag = false", username).First(&user)
	if result.RowsAffected == 0 {
		return true
	}
	return false
}

// Add 添加用户
func (u User) Add(user model.User) bool {
	user.SetCreateInfo()
	result := db.Create(&user)
	return result.RowsAffected == 1
}
