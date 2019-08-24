package service

import (
	"code.qlteam.com/model"
	"github.com/willerhe/webbase/orm"
)

var User user

type user int

func (user) List(user *model.User, page *model.Page, users *[]model.User) error {
	orm.DB.SqlSession.Find(users)
	return nil
}
