package rest

import (
	"code.qlteam.com/model"
	"code.qlteam.com/service"
	"github.com/gin-gonic/gin"
)

func RegisterProjectAPI(r *gin.RouterGroup) {
	p := r.Group("")
	p.GET("/projects", list)
	p.GET("/project/:id", get)
}

// 获取列表
func list(c *gin.Context) {
	projects := &[]model.Project{}
	service.Project.SqlSession.Find(projects)
	c.JSON(200, projects)
}

// 获取单个
func get(c *gin.Context) {
}
