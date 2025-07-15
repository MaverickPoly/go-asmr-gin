package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func HandleCalculate(c *gin.Context) {
	weightStr := c.Query("weight") // kg
	heightStr := c.Query("height") // meters

	weight, weightOk := strconv.ParseFloat(weightStr, 64)
	height, heightOk := strconv.ParseFloat(heightStr, 64)

	if weightOk != nil || heightOk != nil {
		c.JSON(400, gin.H{"error": "Invalid weight or height!"})
		return
	}

	// Result
	bmi := weight / (height * height)
	status := "Normal"

	if bmi < 20 {
		status = "Underweight"
	}
	if bmi > 25 {
		status = "Overweight"
	}

	c.JSON(200, gin.H{
		"weight": weight,
		"height": height,
		"result": bmi,
		"status": status,
	})
}

func main() {
	app := gin.Default()
	app.GET("/calculate", HandleCalculate)

	app.Run(":3000")
}
