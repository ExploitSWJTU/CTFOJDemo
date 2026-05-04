package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, ok := c.Get("role")
		if !ok || role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{
				"code": 403,
				"msg":  "Admin access required",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
