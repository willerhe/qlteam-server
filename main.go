package main

import (
	"code.qlteam.com/model"
	"code.qlteam.com/router"
	"github.com/willerhe/webbase/apper"
)

func main() {
	app := apper.New()

	// 迁移模型
	model.Sync()
	router.Mount(&app.RouterGroup)

	defer app.Start()

}
