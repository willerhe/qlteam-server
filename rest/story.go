package rest

import (
	"code.qlteam.com/model"
	"code.qlteam.com/service"
	"github.com/gin-gonic/gin"
)

type Story int

func (Story) list(c *gin.Context) {
	stories := &[]model.Story{}
	service.Story.List(stories)
	c.JSON(200, stories)
}

func (Story) get(c *gin.Context) {
	s := &model.Story{}
	service.Story.Get(s)
	c.JSON(200, s)
}

func (s *Story) Register(r *gin.RouterGroup) {
	st := r.Group("")
	st.GET("/stories", s.list)
	st.POST("/stories", s.get)

}
