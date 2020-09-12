package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/llh4github/go-admin-api/exp"
	"github.com/llh4github/go-admin-api/global"
	"github.com/llh4github/go-admin-api/service"
	"github.com/llh4github/go-admin-api/utils"
	"github.com/llh4github/go-admin-api/vo"
)

// CasbinAuth Casbin 鉴权中间件
func CasbinAuth(c *gin.Context) {
	url := c.Request.URL.RequestURI()
	action := c.Request.Method
	jwt := c.GetHeader("auth")
	rns, size := utils.GetRoleNames(jwt)
	if size == 0 {
		global.MyLog.Debugf("No Role Info for URL(%s).", url)
	}
	rule := vo.AuthRule{
		URL:       url,
		Action:    action,
		RoleNames: rns,
	}
	has := service.HasPermission(rule)
	if has {
		c.Next()
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": exp.NotPermission,
			"msg":  "没有权限",
			"data": nil,
		})
		c.Abort()
	}

}
