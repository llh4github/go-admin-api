//+build wireinject

package wirex

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/llh4github/go-admin-api/api"
	"github.com/llh4github/go-admin-api/global"
	"github.com/llh4github/go-admin-api/model"
	"github.com/llh4github/go-admin-api/service"
	"github.com/llh4github/go-admin-api/utils"
	"github.com/sirupsen/logrus"
)

var commonInits = wire.NewSet(global.InitViper, global.InitLog)
var routerSet = wire.NewSet(global.InitGin, api.LoadAPI)

// InitBase 初始化基本组件
// 方便调用
// 调用顺序很重要
func InitBase() {
	initCommonSet()
	initRouter()
	initDB()
	initCasbin()

	initJwtConf()
	initService()
}
func initService() *service.Base {

	wire.Build(service.InitService)
	return nil
}
func initJwtConf() global.JwtConfig {

	wire.Build(utils.InitJwtConf)
	return global.JwtConfig{}
}
func initCasbin() *casbin.Enforcer {
	wire.Build(global.InitCasbin)
	return nil
}
func initDB() model.DBMigrateError {
	wire.Build(global.InitDB, model.DBMigrate)
	return nil
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
