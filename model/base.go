package model

import (
	"time"

	"github.com/llh4github/go-admin-api/utils"
)

// Base 数据模型的公共字段
type Base struct {
	ID string `gorm:"primaryKey;type:varchar(25)" json:"id"`
	// 在创建时，如果该字段值为零值，则使用当前时间填充
	CreatedAt time.Time `json:"created_at" `
	// 可为空
	UpdatedAt *time.Time `json:"updated_at"`
	// 软删除标识
	RemoveFlag bool `json:"remove_flag"`

	CreatedBy string `json:"created_by" gorm:"type:varchar(25)"`
	UpdatedBy string `json:"updated_by" gorm:"type:varchar(25)"`
}

// SetCreateInfo 创建时的审计信息
func (b *Base) SetCreateInfo() {

	b.ID = utils.NextIDDefalut()
	b.CreatedAt = time.Now()
	b.RemoveFlag = false
	// TODO 完善创建人信息
	b.CreatedBy = "1"
}

// SetUpdateInfo 更新时的审计信息
func (b *Base) SetUpdateInfo() {
	now := time.Now()
	b.UpdatedAt = &now
	// TODO 完善更新人信息
	b.UpdatedBy = "2"
}

// SetRemove 设置移除信息
func (b *Base) SetRemove() {
	b.RemoveFlag = true
}
