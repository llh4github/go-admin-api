package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/llh4github/go-admin-api/exp"
	"github.com/llh4github/go-admin-api/global"
	"github.com/llh4github/go-admin-api/middleware"
	"github.com/llh4github/go-admin-api/vo"
	"github.com/sirupsen/logrus"
)

var (
	api *gin.RouterGroup
	log *logrus.Logger
)

// LoadAPI 加载API
// 返回值没什么实际意义，wire 框架需要
// 所有API需要在这里注册
func LoadAPI(router *gin.Engine) *gin.RouterGroup {
	api = router.Group(global.Conf.APIPrefix)
	api.Use(middleware.HandleWebException)
	api.Use(middleware.CasbinAuth)
	api.GET("", func(c *gin.Context) {
		c.JSON(200, gin.H{"data": "hello"})
	})
	api.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{"data": "test"})
	})
	log = global.MyLog
	userAPI()
	roleAPI()
	accountAPI()
	return api
}

// baseAPI 基础API结构体。
// 	虽然可以不使用结构体，但为了使用公共方法和避免函数名冲突，
// 	所以选用使用结构体来管理各类API处理方法。
type baseAPI struct {
}

// deserialization 以json的形式将请求数据绑定到结构上
// 	model 必须传入指针
func deserialization(c *gin.Context, model interface{}) {
	// BindJSON 方法绑定错误时会在响应Headers写入400的状态码
	if err := c.ShouldBind(model); err != nil {
		log.Errorf("数据绑定错误, %v \n", err)
		panic(exp.GetCommonExp(exp.BindJSONError))
	}
}

// 返回指定格式的json数据
func (baseAPI) respJSON(c *gin.Context, data vo.JSONWrapper) {
	c.JSON(http.StatusOK, data)
}
