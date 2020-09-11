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
	api.POST("/account/login", r.Login)

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

func (a account) Login(c *gin.Context) {
	var acc vo.Account
	deserialization(c, &acc)
	login := a.s.Login(acc)
	a.respJSON(c, vo.OkResponse(login))
}
