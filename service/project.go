package service

import (
	"code.qlteam.com/model"
	"github.com/willerhe/webbase/orm"
)

// 对外使用
var Project project

type project int

func (project) List(projects *[]model.Project) {
	orm.DB.SqlSession.Find(projects)
}
