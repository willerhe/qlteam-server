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

func (s *Story) Register(r *gin.RouterGroup) {
	st := r.Group("")
	st.GET("/stories", s.list)

}
