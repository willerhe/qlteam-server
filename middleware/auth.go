package middleware

import (
	"github.com/gin-gonic/gin"
)

// AllowCORS 跨域
func AllowCORS(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "localhost")
	c.Writer.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	c.Writer.Header().Set("content-type", "application/json")
	c.Next()

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
