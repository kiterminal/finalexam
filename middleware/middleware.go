package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Authorization(c *gin.Context) {
	token := c.GetHeader("authorization")
	if token != "token2019" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You don't have authorized"})
		c.Abort()
		return
	}

	c.Next()
}
