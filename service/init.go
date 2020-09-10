package service

import (
	"github.com/llh4github/go-admin-api/global"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var db *gorm.DB
var log *logrus.Logger

// Base 基础服务
type Base struct {
}

// func init() {
// 	db = global.MyDB
// 	log = global.MyLog
// }

// InitService for wire tool
func InitService() *Base {
	db = global.MyDB
	log = global.MyLog
	return nil
}
