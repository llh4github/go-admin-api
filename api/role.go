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
	api.PUT("/role/update", r.Update)
	api.DELETE("/role/delete/:id", r.Delete)
}

type role struct {
	baseAPI
	s service.Role
}

func (r role) Add(c *gin.Context) {
	var role model.Role
	if err := c.BindJSON(&role); err != nil {
		log.Errorf("数据绑定错误, %v \n", err)
	}
	add := r.s.Save(role)
	r.respJSON(c, vo.OkResponse(add))
}

func (r role) Update(c *gin.Context) {
	var role model.Role
	if err := c.BindJSON(&role); err != nil {
		log.Errorf("数据绑定错误, %v \n", err)
	}
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
