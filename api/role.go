package api

import (
	"github.com/gin-gonic/gin"
	"github.com/llh4github/go-admin-api/model"
	"github.com/llh4github/go-admin-api/service"
	"github.com/llh4github/go-admin-api/vo"
)

func roleAPI() {

	r := role{s: service.Role{}}
	api.POST("/role", r.Add)
	api.GET("/role/all", r.All)
	api.GET("/role/user/:user_id", r.FindByUserID)
	api.PUT("/role/update", r.Update)
	api.DELETE("/role/delete/:id", r.Delete)
}

type role struct {
	baseAPI
	s service.Role
}

func (r role) Add(c *gin.Context) {
	var role model.Role
	deserialization(c, &role)
	add := r.s.Save(role)
	r.respJSON(c, vo.OkResponse(add))
}

func (r role) Update(c *gin.Context) {
	var role model.Role
	deserialization(c, &role)
	updated := r.s.Update(role)
	r.respJSON(c, vo.OkResponse(updated))
}

func (r role) Delete(c *gin.Context) {
	id := c.Param("id")
	deleted := r.s.Remove(id)
	r.respJSON(c, vo.OkResponse(deleted))
}
func (r role) All(c *gin.Context) {
	l := r.s.All()
	r.respJSON(c, vo.OkResponse(l))
}

// FindByUserID 根据用户id查找对应角色信息
func (r role) FindByUserID(c *gin.Context) {
	userID := c.Param("user_id")

	l := r.s.FindByUserID(userID)

	r.respJSON(c, vo.OkResponse(l))
}
