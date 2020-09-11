package service

import (
	"github.com/llh4github/go-admin-api/utils"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	db       *gorm.DB
	log      *logrus.Logger
	idWorker *utils.SnowflakeIdWorker
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
