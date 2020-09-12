package service

import (
	"github.com/llh4github/go-admin-api/exp"
	"github.com/llh4github/go-admin-api/model"
	"github.com/llh4github/go-admin-api/utils"
	"github.com/llh4github/go-admin-api/vo"
)

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

// FindByUsername 根据用户名查找用户
func (u User) FindByUsername(username string) model.User {
	var user model.User
	result := db.Where("username = ? and remove_flag = false", username).First(&user)
	if result.RowsAffected == 0 {
		panic(exp.GetCommonExp(exp.DataNotFound))
	}
	return user
}

// UpdateRoles 更新用户与角色的关系
func (u User) UpdateRoles(relation vo.UserRoles) int {

	// 1. 删除之前的关系
	result := db.Where("user_id = ? ", relation.UserID).Delete(&model.UserRole{})
	if len(relation.RoleIDs) == 0 {
		log.Debugf("user(id:%s) remove all roles!", relation.UserID)
		return int(result.RowsAffected)
	}
	// 2. 保存新的关系
	var rl []model.UserRole
	for _, rID := range relation.RoleIDs {
		rl = append(rl, model.UserRole{
			ID:     utils.NextIDDefalut(),
			RoleID: rID,
			UserID: relation.UserID,
		})
	}
	result = db.Create(rl)
	return int(result.RowsAffected)
}
