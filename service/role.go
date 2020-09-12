package service

import (
	"time"

	"github.com/llh4github/go-admin-api/exp"
	"github.com/llh4github/go-admin-api/model"
)

// Role 操作角色信息
type Role struct {
}

func (Role) findByID(id string) model.Role {
	var found model.Role
	result := db.Where("id = ? and remove_flag = false", id).First(&found)
	if result.RowsAffected != 1 {
		log.Error(result.Error)
		panic(exp.GetCommonExp(exp.DataNotFound))
	}
	return found
}

// Save 添加角色信息
func (Role) Save(model model.Role) bool {
	model.SetCreateInfo()
	result := db.Create(&model)
	return result.RowsAffected == 1
}

// Update 更新
func (r Role) Update(mdl model.Role) int {
	found := r.findByID(mdl.ID)
	m := map[string]interface{}{
		UpdatedBy:      2,
		UpdatedAt:      time.Now(),
		"role_name":    mdl.RoleName,
		"remark":       mdl.Remark,
		"display_name": mdl.DisplayName,
	}

	result := db.Model(&found).Updates(m)
	return int(result.RowsAffected)
}

// Remove 软删除
func (r Role) Remove(id string) int {
	found := r.findByID(id)

	m := map[string]interface{}{
		UpdatedBy:  3,
		UpdatedAt:  time.Now(),
		RemoveFlag: true,
	}
	result := db.Model(&found).Updates(m)
	return int(result.RowsAffected)
}

// All 查询所有（仅测试用）
func (Role) All() []model.Role {
	var list []model.Role
	db.Where("remove_flag = false").Find(&list)
	return list
}

// FindByUserID 根据用户id查找对应角色信息
func (r Role) FindByUserID(userID string) (roles []model.Role) {

	var rl []model.UserRole
	db.Where("user_id = ?", userID).Find(&rl)
	if len(rl) == 0 {
		return
	}

	var rIDs []string
	for _, ele := range rl {
		rIDs = append(rIDs, ele.RoleID)
	}
	db.Where("id in ?", rIDs).Find(&roles)

	return
}
