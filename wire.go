//+build wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/llh4github/go-admin-api/api"
	"github.com/llh4github/go-admin-api/global"
	"github.com/llh4github/go-admin-api/model"
	"github.com/sirupsen/logrus"
)

var commonInits = wire.NewSet(global.InitViper, global.InitLog)
var routerSet = wire.NewSet(global.InitGin, api.LoadAPI)

func init() {
	initBase()
}

// InitBase 初始化基本组件
// 方便调用
// 调用顺序很重要
func initBase() {
	initAppPath()
	initCommonSet()
	initRouter()
	initDB()
}
func initDB() model.DBMigrateError {
	wire.Build(global.InitDB, model.DBMigrate)
	return nil
}
func initAppPath() appLocal {
	wire.Build(appPath)
	return ""
}

// initCommonSet 初始化通用组件
func initCommonSet() *logrus.Logger {
	wire.Build(commonInits)
	return nil
}

// initRouter gin组件
func initRouter() *gin.RouterGroup {
	wire.Build(routerSet)
	return nil
}
