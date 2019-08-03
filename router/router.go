package router

import (
	"code.qlteam.com/rest"
	"github.com/gin-gonic/gin"
	"github.com/willerhe/webbase/apper"
)

func Mount(root *gin.RouterGroup) {
	api := root.Group("api")
	register(api, new(rest.Project)) // 项目api
}

func register(root *gin.RouterGroup, child apper.IRouter) {
	child.Register(root)
}
