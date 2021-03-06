package service

import (
	"github.com/casbin/casbin/v2"
	"github.com/llh4github/go-admin-api/global"
	"github.com/llh4github/go-admin-api/utils"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	db       *gorm.DB
	log      *logrus.Logger
	idWorker *utils.SnowflakeIdWorker
	enforcer *casbin.Enforcer
)

const (
	ID         = "id"
	CreatedAt  = "created_at"
	UpdatedAt  = "updated_at"
	RemoveFlag = "remove_flag"
	CreatedBy  = "created_by"
	UpdatedBy  = "updated_by"
)

// Base 基础服务
type Base struct {
}

// InitService for wire tool
func InitService() *Base {
	db = global.MyDB
	log = global.MyLog
	enforcer = global.Enforcer
	return nil
}
