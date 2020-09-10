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
}

type role struct {
	baseAPI
	s service.Role
}

func (r role) Add(c *gin.Context) {
	// s := service.Role{}
	var role model.Role
	if err := c.BindJSON(&role); err != nil {
		log.Errorf("数据绑定错误, %v \n", err)
	}
	add := r.s.Save(role)
	r.respJSON(c, vo.OkResponse(add))
}
