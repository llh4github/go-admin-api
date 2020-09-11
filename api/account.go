package api

import (
	"github.com/gin-gonic/gin"
	"github.com/llh4github/go-admin-api/service"
	"github.com/llh4github/go-admin-api/vo"
)

type account struct {
	baseAPI
	s service.Account
}

func accountAPI() {
	r := account{s: service.NewAccountService()}
	api.POST("/account/register", r.Register)
	// api.GET("/role/all", r.All)
	// api.PUT("/role/update", r.Update)
	// api.DELETE("/role/delete/:id", r.Delete)
}
func (a account) Register(c *gin.Context) {

	var acc vo.Account
	deserialization(c, &acc)
	isUnique := service.UniqueUsername(acc.Username)
	if !isUnique {
		panic("用户名已存在！")
	}
	register := a.s.RegisterAccount(acc)
	a.respJSON(c, vo.OkResponse(register))
}
