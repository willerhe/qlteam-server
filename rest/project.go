package rest

import (
	"code.qlteam.com/model"
	"code.qlteam.com/service"
	"github.com/gin-gonic/gin"
)

type Project int

func (p *Project) Register(router *gin.RouterGroup) {
	r := router.Group("")
	r.GET("/projects", p.list)
	r.GET("/project/:id", p.get)
}

// 获取列表
func (*Project) list(c *gin.Context) {

	projects := &[]model.Project{}
	service.Project.List(projects)

	c.JSON(200, projects)
}

// 获取单个
func (*Project) get(c *gin.Context) {
}
