package model

import (
	"github.com/willerhe/webbase/app"
	"github.com/willerhe/webbase/model"
	"time"
)

type (
	// xp 不需要branch
	Project struct {
		model.Base
		Name     string `json:"name"`    // 项目名称
		Describe string `json:"content"` // 项目描述
	}

	// Task 任务
	// 1 作为开发人员，我可以随时查看自己的收件箱有没有邮件，以便能够及时收到领导的任务分配
	// 2 作为项目经理，我可以把用户故事分配给每个开发人员，以便能够按时完成用户需求
	Task struct {
		model.Base
		Name     string    `json:"name"`
		Describe string    `json:"describe"`
		DeadLine time.Time `json:"deadLine"` // 最后期限
		Status   string    `json:"status"`   // 任务状态 未开始  进行中 已完成
		Box      string    `json:"box"`      // 在哪个箱子 收件箱 今天做  下一步做 以后再做
		Remark   string    `json:"remark"`   // 备注
	}

	// 用户故事简单到没有名称 直接是一段对开发工作和用户价值对应关系的描述
	Story struct {
		model.Base
		Describe string `json:"describe"`
	}
)

func Sync() {
	sync := app.NewService()
	sync.SqlSession.AutoMigrate(new(Project), new(Task), new(Story))
}
