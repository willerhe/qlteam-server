package middleware

import "github.com/gin-gonic/gin"

// 允许跨域
func AllowCORS(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	c.Writer.Header().Set("content-type", "application/json")
	c.Next()
}

// 验证当前用户是否登录
func MustLogged(c *gin.Context) {
	// todo 验证是否登录
	c.Next()
}
