package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/llh4github/go-admin-api/global"
	"github.com/llh4github/go-admin-api/vo"
)

var (
	api *gin.RouterGroup
	log = global.MyLog
)

// LoadAPI 加载API
// 返回值没什么实际意义，wire 框架需要
// 所有API需要在这里注册
func LoadAPI(router *gin.Engine) *gin.RouterGroup {
	api = router.Group(global.Conf.APIPrefix)
	api.GET("", func(c *gin.Context) {
		c.JSON(200, gin.H{"data": "hello"})
	})
	api.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{"data": "test"})
	})
	userAPI()
	roleAPI()
	return api
}

// baseAPI 基础API结构体。
// 	虽然可以不使用结构体，但为了使用公共方法和避免函数名冲突，
// 	所以选用使用结构体来管理各类API处理方法。
type baseAPI struct {
}

// 返回指定格式的json数据
func (baseAPI) respJSON(c *gin.Context, data vo.JSONWrapper) {
	c.JSON(http.StatusOK, data)
}
