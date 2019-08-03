package rest

import (
	"code.qlteam.com/model"
	"code.qlteam.com/service"
	"github.com/gin-gonic/gin"
)

type Task int

func (Task) list(c *gin.Context) {
	tasks := &[]model.Task{}
	service.Task.List(tasks)

	c.JSON(200, tasks)
}

func (s *Task) Register(r *gin.RouterGroup) {
	st := r.Group("")
	st.GET("/tasks", s.list)

}
