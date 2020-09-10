package service

import "github.com/llh4github/go-admin-api/model"

// Role 操作角色信息
type Role struct {
}

// Save 添加角色信息
func (Role) Save(model model.Role) bool {
	log.Debug("model: ", model)
	model.SetCreateInfo()

	result := db.Create(&model)
	return result.RowsAffected == 1
}
