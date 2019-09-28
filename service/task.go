package service

import (
	"code.qlteam.com/dict"
	"code.qlteam.com/model"
	"fmt"
	"github.com/willerhe/webbase/orm"
	"log"
	"time"
)

var Task task

type task int

// List 查询所有的自己的task
func (task) List(form *model.Task, ts *[]model.Task) bool {
	form.DeletedAt = nil
	if err := orm.DB.SqlSession.Where(form).Where("on_file = ?", form.OnFile).Find(ts).Error; err != nil {
		log.Println(err)
		return false
	}
	// 查询成功
	return true
}

// Create create  a new task
func (task) Create(t *model.Task, user model.User) bool {
	t.Creator = user.ID
	t.Status = dict.ProjectStatus_Preparing
	if t.Kind == "private" {
		privateTask(t, user)
	}
	if err := orm.DB.SqlSession.Omit("updated_at", "deleted_at", "dead_line").Create(t).Error; err != nil {
		log.Println(err)
		return false
	}
	// 创建成功
	return true
}

// Delete 删除item
func (task) Delete(t *model.Task, user model.User) bool {
	if err := orm.DB.SqlSession.Where(t).Delete(t).Error; err != nil {
		log.Println(err)
		return false
	}
	return true
}

// Update 更新
func (task) Update(t *model.Task, user model.User) bool {
	// todo 根据属性动态选择更新的字段
	if err := orm.DB.SqlSession.Model(t).Updates(*t).Error; err != nil {
		log.Println(err)
		return false
	}
	return true
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
	case dict.Box_TODO:
		deadline = time.Now().AddDate(0, 0, 1)
	case dict.BOX_NextStep:
		deadline = time.Now().AddDate(0, 0, 3)
	case dict.Box_Later:
		deadline = time.Now().AddDate(0, 0, 7)
	case dict.Box_Inbox:
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
