package service

import (
	"code.qlteam.com/errset"
	"code.qlteam.com/model"
	"github.com/willerhe/webbase/orm"
)

type login int

var Login login

// Register 注册用户
func (login) Register(u *model.User) error {
	// 1. 判断是否存在

	count := orm.DB.SqlSession.Where("account = ?", u.Account).First(u).RowsAffected
	if count > 0 {
		return errset.AccountExisted
	}
	// 2. 新建用户
	orm.DB.SqlSession.NewRecord(u)
	orm.DB.SqlSession.Omit("deleted_at", "updated_at").Create(u)

	return nil
}
