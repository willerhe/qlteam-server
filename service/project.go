package service

import (
	"code.qlteam.com/model"
	"github.com/willerhe/webbase/orm"
)

// 组织模块
type project int

// 对外使用
var Project project

var db = orm.NewService()

func (p *project) List(projects *[]model.Project) {
	db.SqlSession.Find(projects)
}
