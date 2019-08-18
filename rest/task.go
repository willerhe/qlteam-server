package rest

import (
	"code.qlteam.com/model"
	"code.qlteam.com/service"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

type Task int

// list 查询tasks
func (Task) list(c *gin.Context) {
	form := &model.Task{}
	if err := c.Bind(form); err != nil {
		log.Println(err)
		c.String(400, "参数错误")
		c.Abort()
		return
	}
	u, _ := c.Get("user")
	form.Leader = u.(model.User).ID

	tasks := &[]model.Task{}
	if !service.Task.List(form, tasks) {
		c.String(500, "查询失败")
		c.Abort()
		return
	}

	c.JSON(200, tasks)
}

// create 创建新task
func (Task) create(c *gin.Context) {
	form := new(model.Task)
	if err := c.Bind(form); err != nil {
		log.Println(err)
		c.String(200, "参数绑定错误")
		c.Abort()
		return
	}
	u, _ := c.Get("user")
	if !service.Task.Create(form, u.(model.User)) {
		c.String(500, "常见新任务失败")
		c.Abort()
		return
	}

	c.JSON(200, form)
}

// update update a task
func (Task) update(c *gin.Context) {
	form := new(model.Task)
	if err := c.Bind(form); err != nil {
		log.Println(err)
		c.String(400, "请检查更新参数")
		c.Abort()
		return
	}

	u, _ := c.Get("user")
	if !service.Task.Update(form, u.(model.User)) {
		c.String(500, "更新task失败")
		c.Abort()
		return
	}
	c.JSON(200, form)
}

// delete delete a task
func (Task) delete(c *gin.Context) {
	form := new(model.Task)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.String(400, "请检查更新参数")
		c.Abort()
		return
	}
	form.ID = uint(id)

	u, _ := c.Get("user")
	if !service.Task.Delete(form, u.(model.User)) {
		c.String(500, "更新task失败")
		c.Abort()
		return
	}
	c.JSON(200, form)
}

// Register register a group of router to root router
func (s *Task) Register(r *gin.RouterGroup) {
	st := r.Group("")
	st.GET("/tasks", s.list)
	st.POST("/task", s.create)
	st.PUT("/task/:id", s.update)
	st.DELETE("/task/:id", s.delete)
}
