package rest

import (
	"code.qlteam.com/model"
	"code.qlteam.com/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Story int

// list
func (Story) list(c *gin.Context) {
	stories := &[]model.Story{}
	service.Story.List(stories)
	c.JSON(200, stories)
}

// get
func (Story) get(c *gin.Context) {
	s := &model.Story{}
	id, _ := strconv.Atoi(c.Param("id"))
	s.ID = uint(id)
	service.Story.Get(s)
	c.JSON(200, s)
}

// delete
func (Story) delete(c *gin.Context) {
	s := &model.Story{}
	id, _ := strconv.Atoi(c.Param("id"))
	s.ID = uint(id)
	service.Story.Delete(s)
	c.JSON(200, s)
}

func (Story) create(c *gin.Context) {
	s := &model.Story{}
	c.Bind(s)
	service.Story.Create(s)
	c.JSON(200, s)
}

// save
func (Story) update(c *gin.Context) {
	s := &model.Story{}
	c.Bind(s)
	service.Story.Update(s)
	c.JSON(200, s)
}

//Register register api
func (s *Story) Register(r *gin.RouterGroup) {
	st := r.Group("")
	st.GET("/story/:id", s.get)
	st.GET("/stories", s.list)
	st.POST("/story", s.create)
	st.PUT("/story", s.update)
	st.DELETE("/story/:id", s.delete)

}
