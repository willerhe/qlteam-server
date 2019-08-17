package middleware

import (
	"code.qlteam.com/model"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/willerhe/webbase/configer"
)

// AllowCORS 跨域
func AllowCORS() func(c *gin.Context) {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("origin")
		c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		c.Writer.Header().Add("Access-Control-Allow-Headers", "Content-Type")
		c.Writer.Header().Set("content-type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Accept,authorization,Authorization,Cache-Control,Content-Type,DNT,If-Modified-Since,Keep-Alive,Origin,User-Agent,X-Mx-ReqToken,X-Requested-With")
		if method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}

}

// MustLogged 必须登录
func MustLogged(c *gin.Context) {

	t := c.GetHeader("authorization")
	if t == "" {
		c.String(401, "请登录后再进行操作！")
		c.Abort()
		return
	}
	//todo TOKEN3 验证token
	token, err := jwt.ParseWithClaims(t, &model.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(configer.Config.Get("token.sign").(string)), nil
	})

	if claims, ok := token.Claims.(*model.Claims); ok && token.Valid {
		fmt.Println("当前用户", claims.User)
		c.Set("user", claims.User)
	} else {
		c.String(401, "登录信息过期请重新登录")
		c.Abort()
		fmt.Println(err)
		return
	}
	c.Next()
}
