package api

import (
	"github.com/gin-gonic/gin"
	"github.com/llh4github/go-admin-api/model"
	"github.com/llh4github/go-admin-api/service"
	"github.com/llh4github/go-admin-api/vo"
)

func userAPI() {
	r := user{s: service.User{}}
	api.GET("/user", r.GET)
	api.POST("/user", r.AddUser)
	api.PUT("/user/update/role", r.UpdateRole)
}

type user struct {
	baseAPI
	s service.User
}

func (m *user) GET(c *gin.Context) {
	r := "hello"
	m.respJSON(c, vo.OkResponse(r))
}
func (m *user) AddUser(c *gin.Context) {
	var user model.User
	deserialization(c, &user)
	m.respJSON(c, vo.OkResponse(user))
}
func (m *user) UpdateRole(c *gin.Context) {
	var relation vo.UserRoles
	deserialization(c, &relation)
	updated := m.s.UpdateRoles(relation)
	m.respJSON(c, vo.OkResponse(updated))
}
