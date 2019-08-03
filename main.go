package main

import (
	"code.qlteam.com/model"
	"code.qlteam.com/router"
	"github.com/willerhe/webbase/app"
)

func main() {
	app.Load()
	// 迁移模型
	model.Sync()
	router.Mount(app.Route)

	defer app.Run()

}
