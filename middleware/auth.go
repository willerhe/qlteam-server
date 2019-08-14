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
		c.String(401, "请先登录！")
		c.Abort()
		return
	}
	//todo TOKEN3 验证token
	token, err := jwt.ParseWithClaims(t, &model.Token{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(configer.Config.Get("token.sign").(string)), nil
	})

	if claims, ok := token.Claims.(*model.Token); ok && token.Valid {
		fmt.Println("当前用户", claims.UID)
	} else {
		fmt.Println(err)
	}

	c.Next()
}
