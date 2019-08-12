package middleware

import (
	"github.com/gin-gonic/gin"
)

// AllowCORS 跨域
func AllowCORS() func(c *gin.Context) {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("origin")

		if method == "OPTIONS" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
			c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			c.Writer.Header().Set("Access-Control-Allow-Headers", "Accept,Authorization,Cache-Control,Content-Type,DNT,If-Modified-Since,Keep-Alive,Origin,User-Agent,X-Mx-ReqToken,X-Requested-With")
			c.AbortWithStatus(204)
			return
		}

		c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Next()
		return
	}

}

// MustLogged 必须登录
func MustLogged(c *gin.Context) {
	token := c.GetHeader("authorization")
	if token == "" {
		c.String(401, "请先登录！")
		c.Abort()
		return
	}
	c.Next()
}
