package service

import (
	"code.qlteam.com/model"
	"github.com/willerhe/webbase/orm"
)

var Story story

type story int

// List story list
func (story) List(ss *[]model.Story) {
	orm.DB.SqlSession.Find(ss)
}

// Get get single item
func (story) Get(s *model.Story) {
	orm.DB.SqlSession.First(s)
}

// Save save new story
func (story) Save(s *model.Story) {
	orm.DB.SqlSession.Save(s)
}

func (story) Delete(s model.Story) {
	orm.DB.SqlSession.Delete(s)
}
