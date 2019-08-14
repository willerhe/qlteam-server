package service

import (
	"code.qlteam.com/model"
	"github.com/willerhe/webbase/orm"
)

var Task task

type task int

func (task) List(ts *[]model.Task) {
	orm.DB.SqlSession.Find(ts)
}

func (task) Create(t *model.Task) {
	orm.DB.SqlSession.Insert(t)
}
