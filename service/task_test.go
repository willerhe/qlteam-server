package service

import (
	"code.qlteam.com/model"
	"fmt"
	"testing"
	"time"
)

func TestLogin_Register(test *testing.T) {
	//timeStr := time.Now().Format("2006-01-02")
	//fmt.Println("timeStr:", timeStr)
	//t, _ := time.Parse("2006-01-02", timeStr)
	//timeNumber := t.Unix()
	//fmt.Println(t,"timeNumber:", timeNumber)

	zeroTime, _ := time.Parse("2006-01-02", time.Now().Format("2006-01-02"))
	fmt.Println(zeroTime)
	fmt.Println(time.Now().Format("060102"))

	fmt.Println("------")
	u := model.User{}
	u.ID = 1
	fmt.Println(generalTaskName(u))
}
