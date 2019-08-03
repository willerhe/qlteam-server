package service

import (
	"code.qlteam.com/model"
	"github.com/willerhe/webbase/orm"
)

var Story story

type story int

func (story) List(s *[]model.Story) {
	orm.DB.SqlSession.Find(s)
}
