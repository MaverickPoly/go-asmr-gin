package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

var exchangeRates = map[string]float64{
	"USD": 1.0,
	"EUR": 0.85,
	"RUB": 93.5,
	"CNY": 7.25,
}

func ConvertHandler(c *gin.Context) {
	from := c.Query("from")
	to := c.Query("to")
	amountStr := c.DefaultQuery("amount", "1")

	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil || amount <= 0 {
		c.JSON(400, gin.H{"error": "Invalid amount!"})
		return
	}

	fromRate, fromExists := exchangeRates[from]
	toRate, toExists := exchangeRates[to]

	if !fromExists || !toExists {
		c.JSON(400, gin.H{"error": "Unsupported currency code!"})
		return
	}

	usdAmount := amount / fromRate
	converted := usdAmount * toRate

	c.JSON(200, gin.H{
		"from":     from,
		"to":       to,
		"amount":   amount,
		"result":   converted,
		"exchange": toRate / fromRate,
	})

}

func main() {
	app := gin.Default()

	app.GET("/convert", ConvertHandler)

	app.Run(":8000")
}
