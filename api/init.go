package api

import (
	"github.com/gin-gonic/gin"
	"github.com/llh4github/go-admin-api/global"
)

// LoadAPI 加载API
// 返回值没什么实际意义，wire 框架需要
// 所有API需要在这里注册
func LoadAPI(router *gin.Engine) *gin.RouterGroup {
	api := router.Group(global.Conf.APIPrefix)
	api.GET("", func(c *gin.Context) {
		c.JSON(200, gin.H{"data": "hello"})
	})
	api.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{"data": "test"})
	})
	return api
}
