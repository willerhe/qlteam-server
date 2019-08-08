package rest

import (
	"code.qlteam.com/model"
	"code.qlteam.com/service"
	"fmt"
	"testing"
)

func TestStory(t *testing.T) {
	s := &model.Story{
		Describe: "测试使用的用户故事",
	}
	service.Story.Create(s)
	fmt.Println("创建后的id是", s.ID)
	s.Describe = ""
	service.Story.Update(s)
	fmt.Println("修改用户故事", s)
	service.Story.Delete(s)
	fmt.Println("删除用户故事，id 为", s.ID)
}
