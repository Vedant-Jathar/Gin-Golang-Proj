package middleware

import "github.com/gin-gonic/gin"

func AuthMidllware(c *gin.Context) {
	c.Set("userId", 1)
	c.Next()
}
