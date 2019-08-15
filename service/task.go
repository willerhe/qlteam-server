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
	orm.DB.SqlSession.Omit("updated_at", "deleted_at").Create(t)
}

// Delete 删除item
func (task) Delete(t *model.Task) {
	orm.DB.SqlSession.Where(t).Delete(t)
}

// Update 更新
func (task) Update(t *model.Task) {
	orm.DB.SqlSession.Omit("deleted_at").Save(t)
}

// privateTask 私人任务
func privateTask(t *model.Task, user model.User) {
	t.Leader = user.ID
	t.Organizer = user.ID
	//	 todo 根据box 计算deadline 放在constant 中
	// 190815001 2019年8月15日的第一项任务
	t.Name = generalTaskName(user)
	t.Kind = "default"
	// todo deadline 生成后  后期怎么续期
	t.DeadLine = generalDeadline(t.Kind)

}

//inbox
//todo
//nextStep
//later

// 截止日计算
func generalDeadline(s string) time.Time {
	var deadline time.Time
	switch s {
	case "todo":
		deadline = time.Now().AddDate(0, 0, 1)
	case "nextStep":
		deadline = time.Now().AddDate(0, 0, 3)
	case "later":
		deadline = time.Now().AddDate(0, 0, 7)
	case "inbox":
		deadline = time.Now().AddDate(0, 0, 1)

	}
	return deadline
}

// computedTaskName 根据当前的时间计算今天的任务数字
func generalTaskName(user model.User) string {
	var count int
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
