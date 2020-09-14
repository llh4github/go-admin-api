package api

import (
	"github.com/gin-gonic/gin"
	"github.com/llh4github/go-admin-api/service"
	"github.com/llh4github/go-admin-api/vo"
)

type permission struct {
	baseAPI
	s service.CasbinX
}

func permissionAPI() {

	r := permission{s: service.CasbinX{}}
	api.POST("/permission", r.Add)
	api.GET("/permission/all", r.All)
	// api.GET("/role/user/:user_id", r.FindByUserID)
	// api.PUT("/role/update", r.Update)
	api.DELETE("/permission/delete", r.Delete)
}

func (p permission) Add(c *gin.Context) {
	var rule vo.PermInfo
	deserialization(c, &rule)
	add := p.s.Add(rule)
	p.respJSON(c, vo.OkResponse(add))
}

func (p permission) All(c *gin.Context) {

	list := p.s.All()
	p.respJSON(c, vo.OkResponse(list))
}

func (p permission) Delete(c *gin.Context) {
	var rule vo.PermInfo
	deserialization(c, &rule)
	deleted := p.s.Delete(rule)
	p.respJSON(c, vo.OkResponse(deleted))
}
