// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package wirex

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/llh4github/go-admin-api/api"
	"github.com/llh4github/go-admin-api/global"
	"github.com/llh4github/go-admin-api/model"
	"github.com/llh4github/go-admin-api/utils"
	"github.com/sirupsen/logrus"
)

// Injectors from wire.go:

func initJwtConf() global.JwtConfig {
	jwtConfig := utils.InitJwtConf()
	return jwtConfig
}

func initCasbin() *casbin.Enforcer {
	enforcer := global.InitCasbin()
	return enforcer
}

func initDB() model.DBMigrateError {
	db := global.InitDB()
	dbMigrateError := model.DBMigrate(db)
	return dbMigrateError
}

func initCommonSet() *logrus.Logger {
	configuration := global.InitViper()
	logger := global.InitLog(configuration)
	return logger
}

func initRouter() *gin.RouterGroup {
	engine := global.InitGin()
	routerGroup := api.LoadAPI(engine)
	return routerGroup
}

// wire.go:

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
}
