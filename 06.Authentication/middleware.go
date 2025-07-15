package main

import (
	"github.com/gin-gonic/gin"
)

func AuthRequired(c *gin.Context) {
	accessToken, err := c.Cookie("accessToken")
	if err != nil {
		c.AbortWithStatusJSON(401, gin.H{"error": "No access token provided!"})
		return
	}
	claims, err := ValidateToken(accessToken)
	if err != nil {
		c.AbortWithStatusJSON(401, gin.H{"error": "Invalid token!"})
		return
	}
	c.Set("userId", claims.UserId)
	c.Next()
}
