package rest

import (
	"code.qlteam.com/model"
	"code.qlteam.com/service"
	"github.com/gin-gonic/gin"
	"log"
)

type User int

// list 查询tasks
func (User) list(c *gin.Context) {
	form := &model.User{}

	if err := c.Bind(form); err != nil {
		log.Println(err)
		c.String(400, "参数错误")
		c.Abort()
		return
	}
	page := &model.Page{}
	if err := c.Bind(page); err != nil {
		log.Println(err)
		c.String(400, "参数错误")
		c.Abort()
		return
	}
	users := &[]model.User{}
	if err := service.User.List(form, page, users); err != nil {
		c.String(500, "查询失败")
		c.Abort()
		return
	}
	r := make(map[string]interface{}, 2)
	r["data"] = users
	r["page"] = page

	c.JSON(200, r)
}

// Register register a group of router to root router
func (s *User) Register(r *gin.RouterGroup) {
	st := r.Group("")
	st.GET("/users", s.list)
}
