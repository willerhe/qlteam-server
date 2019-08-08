package rest

import "github.com/gin-gonic/gin"

type Login int

func (Login) login(c *gin.Context) {
	c.String(200, "ok")
}

func (l *Login) Register(router *gin.RouterGroup) {
	r := router.Group("")
	r.GET("/projects", l.login)
	//r.GET("/project/:id", p.get)
}
