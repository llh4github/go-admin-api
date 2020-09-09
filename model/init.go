// Package model 模型层只定义数据库表对应的模型，不实现相关操作
//
// 数据库相关操作放到 service 层完成
package model

import (
	"errors"

	"github.com/llh4github/go-admin-api/global"
	"gorm.io/gorm"
)

// DBMigrateError 数据库迁移错误
// 此错误仅用于 wire 工具
type DBMigrateError error

// DBMigrate 模型注册
// 自动进行数据库迁移
func DBMigrate(db *gorm.DB) DBMigrateError {
	err := db.AutoMigrate(&User{})
	if err != nil {
		global.MyLog.Fatalf("数据库迁移失败！ %v ", err)
	}
	return DBMigrateError(errors.New("数据库迁移失败！"))
}
