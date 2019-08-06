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
// todo 更细其他字段 不更新deleted_at字段
func (story) Update(s *model.Story) {
	orm.DB.SqlSession.Model(s).UpdateColumn("describe")
}

// Delete delete story
func (story) Delete(s *model.Story) {
	orm.DB.SqlSession.Delete(s)
}

// Create create new story
func (story) Create(s *model.Story) {
	s = orm.DB.SqlSession.Create(s).Value.(*model.Story)
}
