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
	// 跟路由
	root := &app.RouterGroup
	api := root.Group("api")
	// api 子路由
	v1 := api.Group("v1")

	// 登录
	router.Register(root, new(rest.Login))

	v1.Use(middleware.AllowCORS, middleware.MustLogged)

	router.Register(v1, new(rest.Project)) // 项目api
	router.Register(v1, new(rest.Story))   // 用户故事api
	router.Register(v1, new(rest.Task))    // 任务
	router.Register(v1, new(rest.Login))

}
