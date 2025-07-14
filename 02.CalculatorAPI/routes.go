package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/add", func(c *gin.Context) {
		a, errA := strconv.ParseFloat(c.Query("a"), 64)
		b, errB := strconv.ParseFloat(c.Query("b"), 64)

		if errA != nil || errB != nil {
			c.JSON(http.StatusBadGateway, gin.H{"error": "Invalid query!"})
			return
		}

		res := a + b
		c.JSON(http.StatusOK, gin.H{"result": res})
	})
	r.GET("/subtract", func(c *gin.Context) {
		a, errA := strconv.ParseFloat(c.Query("a"), 64)
		b, errB := strconv.ParseFloat(c.Query("b"), 64)

		if errA != nil || errB != nil {
			c.JSON(http.StatusBadGateway, gin.H{"error": "Invalid query!"})
			return
		}

		res := a - b
		c.JSON(http.StatusOK, gin.H{"result": res})
	})
	r.GET("/multiply", func(c *gin.Context) {
		a, errA := strconv.ParseFloat(c.Query("a"), 64)
		b, errB := strconv.ParseFloat(c.Query("b"), 64)

		if errA != nil || errB != nil {
			c.JSON(http.StatusBadGateway, gin.H{"error": "Invalid query!"})
			return
		}

		res := a * b
		c.JSON(http.StatusOK, gin.H{"result": res})
	})
	r.GET("/divide", func(c *gin.Context) {
		a, errA := strconv.ParseFloat(c.Query("a"), 64)
		b, errB := strconv.ParseFloat(c.Query("b"), 64)

		if errA != nil || errB != nil {
			c.JSON(http.StatusBadGateway, gin.H{"error": "Invalid query!"})
			return
		}

		if b == 0.0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Zero division error!"})
			return
		}

		res := a / b
		c.JSON(http.StatusOK, gin.H{"result": res})
	})
}
