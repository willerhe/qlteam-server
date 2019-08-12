package rest

import (
	"code.qlteam.com/errset"
	"code.qlteam.com/model"
	"code.qlteam.com/service"
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/willerhe/webbase/orm"
)

type Login int

// login 登录
func (Login) login(c *gin.Context) {
	u := &model.User{}
	if err := c.Bind(u); err != nil {
		c.String(400, "不支持的参数类型")
		c.Abort()
		return
	}
	// todo 实现remember me
	hash := md5.New()
	hash.Write([]byte(u.Password))

	user := &model.User{Account: u.Account, Password: fmt.Sprintf("%X", hash.Sum(nil))}
	has := orm.DB.SqlSession.Where(user).First(user).RowsAffected > 0
	if !has {
		c.String(403, "用户名或者密码错误")
		c.Abort()
		return
	}
	// 生成jwt
	c.Header("authorization", service.Token.General(*u))

	c.JSON(200, user)
}

// register 注册
func (Login) register(c *gin.Context) {
	u := &model.User{}
	if err := c.Bind(u); err != nil {
		c.String(401, "请检查注册信息是否合法！")
		c.Abort()
	}

	err := service.Login.Register(u)
	if err == errset.AccountExisted {
		c.String(403, "用户已经存在")
		c.Abort()
		return
	}

	c.String(200, "ok")
}

// logout 注销
func (Login) logout(c *gin.Context) {
	c.SetCookie("uid", "", -1, "", "localhost", true, true)
	//todo  删除后台的用户记录
	c.String(200, "ok")
}

func (l *Login) Register(router *gin.RouterGroup) {
	r := router.Group("")
	r.POST("/login", l.login)
	r.POST("/register", l.register)
	r.POST("/logout", l.logout)
	//r.GET("/project/:id", p.get)
}
