package router

import (
	"code.qlteam.com/rest"
	"github.com/gin-gonic/gin"
	"github.com/willerhe/webbase/app"
)

func Mount(root *gin.RouterGroup) {
	api := app.Route.Group("api")
	rest.RegisterProjectAPI(api)
}
