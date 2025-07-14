package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!")
	})
	r.GET("/not-found", func(c *gin.Context) {
		c.String(http.StatusOK, "404 Not Found!")
	})
}
