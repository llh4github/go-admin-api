package api

import (
	"github.com/gin-gonic/gin"
	"github.com/llh4github/go-admin-api/model"
	"github.com/llh4github/go-admin-api/vo"
)

func userAPI() {
	r := user{}
	api.GET("/user", r.GET)
	api.POST("/user", r.AddUser)
}

type user struct {
	baseAPI
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
