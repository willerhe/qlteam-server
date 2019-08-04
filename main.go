package main

import (
	"code.qlteam.com/middleware"
	"code.qlteam.com/model"
	"code.qlteam.com/rest"
	"github.com/willerhe/webbase/apper"
	"github.com/willerhe/webbase/router"
)

func main() {
	app := apper.New()
	defer app.Start()
	// 迁移模型
	model.Sync()
	// 挂在路由
	api := app.Group("api")
	api.Use(middleware.AllowCORS, middleware.MustLogged)

	router.Register(api, new(rest.Project)) // 项目api
	router.Register(api, new(rest.Story))   // 用户故事api
	router.Register(api, new(rest.Task))    // 任务

}
