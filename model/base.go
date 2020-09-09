package model

import (
	"database/sql"
	"time"
)

// Base 数据模型的公共字段
type Base struct {
	ID uint `gorm:"primaryKey" json:"id"`
	// 在创建时，如果该字段值为零值，则使用当前时间填充
	CreatedAt time.Time `json:"created_at" `
	// 可为空
	UpdatedAt *time.Time `json:"updated_at"`
	// 软删除标识
	RemoveFlag bool `json:"remove_flag"`

	CreatedBy uint `json:"created_by"`
	// 可为空
	UpdatedBy sql.NullInt64 `json:"updated_by"`
}

// SetCreateInfo 创建时的审计信息
func (b *Base) SetCreateInfo() {

	b.CreatedAt = time.Now()
	b.RemoveFlag = false
	// TODO 完善创建人信息
	b.CreatedBy = 1
}

// SetUpdateInfo 更新时的审计信息
func (b *Base) SetUpdateInfo() {
	now := time.Now()
	b.UpdatedAt = &now
	u := sql.NullInt64{}
	// TODO 完善更新人信息
	u.Scan(1)
	b.UpdatedBy = u
}

// SetRemove 设置移除信息
func (b *Base) SetRemove() {
	b.RemoveFlag = true
}
