package model

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/willerhe/webbase/modeler"
	"github.com/willerhe/webbase/orm"
	"time"
)

type (
	// xp 不需要branch
	Project struct {
		modeler.Base
		Name     string `json:"name" form:"name"`        // 项目名称
		Describe string `json:"content" form:"describe"` // 项目描述
	}

	// Task 任务
	// 1 作为开发人员，我可以随时查看自己的收件箱有没有邮件，以便能够及时收到领导的任务分配
	// 2 作为项目经理，我可以把用户故事分配给每个开发人员，以便能够按时完成用户需求
	Task struct {
		modeler.Base
		Name     string    `json:"name" form:"name"`
		Describe string    `json:"describe" form:"describe"`
		DeadLine time.Time `json:"deadLine" form:"deadLine"` // 最后期限
		Status   string    `json:"status" form:"status"`     // 任务状态 未开始preparing  进行中ongoing 已完成done
		Box      string    `json:"box" form:"box"`           // 在哪个箱子 收件箱 今天做  下一步做 以后再做
		Remark   string    `json:"remark" form:"remark"`     // 备注

		Kind string `json:"kind" form:"kind"` // 任务类型  private 私人 project

		Group string `json:"group" form:"group"` // 将相同性质的任务放在一组 项目的任务视图使用

		OnFile bool `json:"onFile" form:"onFile"` // 是否已经归档

		Leader    uint `json:"leader" form:"leader"`       // 该项任务的领导者 (负责人)
		Organizer uint `json:"organizer" form:"organizer"` // 组织者

		ProjectId uint `json:"projectId" form:"projectId"` // 所属项目
		Creator   uint `json:"creator" form:"creator"`     // 创建者
	}

	// 用户故事简单到没有名称 直接是一段对开发工作和用户价值对应关系的描述
	Story struct {
		modeler.Base
		ProjectId int    `json:"projectId" form:"projectId"`
		Describe  string `json:"describe" form:"describe"`
	}

	// 用户
	User struct {
		modeler.Base `xorm:"extends"`
		NickName     string `json:"nickName" form:"nickName"`
		Account      string `json:"account" form:"account"`
		Password     string `json:"-" form:"password"`
	}

	// todo TOKEN1 定义token 结构
	Claims struct {
		jwt.StandardClaims
		User User
	}
)

func Sync() {
	sync := orm.NewService()
	sync.SqlSession.AutoMigrate(new(Project), new(Task), new(Story), new(User))
}
