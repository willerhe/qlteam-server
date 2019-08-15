package service

import (
	"code.qlteam.com/model"
	"fmt"
	"github.com/willerhe/webbase/orm"
	"time"
)

var Task task

type task int

func (task) List(ts *[]model.Task) {
	orm.DB.SqlSession.Find(ts)
}

func (task) Create(t *model.Task, user model.User) {
	t.Creator = user.ID
	t.Status = "preparing"

	if t.Kind == "private" {
		privateTask(t, user)
	}
	orm.DB.SqlSession.Omit("updated_at", "deleted_at").Create(user)
}

// privateTask 私人任务
func privateTask(t *model.Task, user model.User) {
	t.Leader = user.ID
	t.Organizer = user.ID
	//	 todo 根据box 计算deadline 放在constant 中
	// 190815001 2019年8月15日的第一项任务
	t.Name = computedTaskName(user)

}

// computedTaskName 根据当前的时间计算今天的任务数字
func computedTaskName(user model.User) string {
	var count int
	// todo 查询当天自己的任务数量 并生成任务序号

	zeroTime, _ := time.Parse("2006-01-02", time.Now().Format("2006-01-02"))
	orm.DB.SqlSession.Where("leader = ? and  created_at BETWEEN ? AND ?", user.ID, zeroTime, time.Now()).Table("tasks").Count(&count)
	//orm.DB.SqlSession.Where("leader = ? ").Count(&count)
	t := time.Now().Format("060102")
	fu := "000"
	c := fmt.Sprint(count + 1)
	// todo 今天的任务量超过了9999个  （基本不可能）
	fu = fu[len(c):]
	r := t + fu + c
	return r
}
